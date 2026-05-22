package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func branchExists(name string) bool {
	if exec.Command("git", "rev-parse", "--verify", "refs/heads/"+name).Run() == nil {
		return true
	}
	return exec.Command("git", "rev-parse", "--verify", "refs/remotes/origin/"+name).Run() == nil
}

func defaultBranch() (string, error) {
	out, err := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD").Output()
	if err == nil {
		ref := strings.TrimSpace(string(out))
		return strings.TrimPrefix(ref, "refs/remotes/origin/"), nil
	}

	out, err = exec.Command("git", "remote", "show", "origin").Output()
	if err == nil {
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "HEAD branch:") {
				return strings.TrimSpace(strings.TrimPrefix(line, "HEAD branch:")), nil
			}
		}
	}

	for _, name := range []string{"main", "master"} {
		if branchExists(name) {
			return name, nil
		}
	}

	return "", fmt.Errorf("could not determine default branch")
}

func checkoutBranch(name string) ([]byte, error) {
	out, err := exec.Command("git", "checkout", name).CombinedOutput()
	if err == nil {
		return out, nil
	}

	return exec.Command("git", "checkout", "-B", name, "origin/"+name).CombinedOutput()
}

func main() {
	orange := color.RGB(209, 94, 20)
	boldOrange := orange.Add(color.Bold).Add(color.Underline)

	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold).Add(color.Underline)

	cyan := color.New(color.FgCyan)
	boldCyan := cyan.Add(color.Bold).Add(color.Underline)

	magenta := color.New(color.FgMagenta)
	boldMagenta := magenta.Add(color.Bold).Add(color.Underline)

	boldRed := color.New(color.FgRed).Add(color.Bold).Add(color.Underline)

	status, err := exec.Command("git", "status").Output()
	if err != nil {
		boldRed.Printf("error getting git status: %v", err)
		os.Exit(1)
	}

	if !strings.Contains(string(status), "working tree clean") {
		boldRed.Println("Working tree is not clean, commit or stash changes before moving on")
		os.Exit(1)
	}

	boldCyan.Println("Fetch all branches and tags from origin:")
	fetch, err := exec.Command("git", "fetch", "origin", "--tags").CombinedOutput()
	if err != nil {
		boldRed.Printf("error fetching all branches and tags from origin: %v", err)
		os.Exit(1)
	}

	fmt.Println(strings.TrimSpace(string(fetch)))

	branch, err := defaultBranch()
	if err != nil {
		boldRed.Printf("error determining default branch: %v\n", err)
		os.Exit(1)
	}

	boldOrange.Printf("Checking out the %s branch:\n", branch)
	checkout, err := checkoutBranch(branch)
	if err != nil {
		boldRed.Printf("error checking out %s branch: %v", branch, err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(checkout)), "\n")

	if len(lines) > 1 && lines[1] != "" {
		fmt.Println(lines[1])
	}

	boldMagenta.Printf("Pull %s branch from origin:\n", branch)
	pull, err := exec.Command("git", "pull", "origin", branch).CombinedOutput()
	if err != nil {
		boldRed.Printf("error pulling %s branch from origin: %v", branch, err)
		os.Exit(1)
	}

	fmt.Println(strings.TrimSpace(string(pull)))

	boldGreen.Println("All set!")
	fmt.Println("")
}
