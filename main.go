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

	words := strings.Fields(string(status))

	for _, word := range words {
		fmt.Println(word)
	}

	// fmt.Println(string(status))
}