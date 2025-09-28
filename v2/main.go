package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(getVersion())
}

func getVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--always", "--dirty", "--match", "v*")
	output, err := cmd.Output()
	if err != nil {
		return "v0.0.0-unknown"
	}
	return strings.TrimSpace(string(output))
}
