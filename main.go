package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Updating main branch from origin")

    status, err := exec.Command("git", "status").Output()
	if err != nil {
		fmt.Printf("error getting git status: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(status))
}