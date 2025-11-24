package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

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

	boldOrange.Println("Checking out the main branch:")
	checkout, err := exec.Command("git", "checkout", "main").CombinedOutput()
	if err != nil {
		boldRed.Printf("error checking out main branch: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(checkout)), "\n")

	if len(lines) > 1 && lines[1] != "" {
		fmt.Println(lines[1])
	}

	boldCyan.Println("Fetch all branches and tags from origin:")
	fetch, err := exec.Command("git", "fetch", "origin", "--tags").CombinedOutput()
	if err != nil {
		boldRed.Printf("error fetching all branches and tags from origin: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(strings.TrimSpace(string(fetch))))

	boldMagenta.Println("Pull main branch from origin:")
	pull, err := exec.Command("git", "pull", "origin", "main").CombinedOutput()
	if err != nil {
		boldRed.Printf("error pulling main branch from origin: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(strings.TrimSpace(string(pull))))

	boldGreen.Println("All set!")
	fmt.Println("")
}
