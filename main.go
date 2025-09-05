package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Updating main branch from origin")

    status, err := exec.Command("git", "status").Output()
	if err != nil {
		fmt.Printf("error getting git status: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(string(status), "\n")

	for _, line := range lines {
		fmt.Println(line)
	}


	// fmt.Println(string(status))
}