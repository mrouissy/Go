package main

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// ClearScreen clears the terminal screen
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}


func ParseID(s string) (int, error) {
	return strconv.Atoi(s)
}

func ValidateContact(name, phone string) bool {
	return name != "" && phone != ""
}

// ConfirmAction checks if user said yes
func ConfirmAction(response string) bool {
	r := strings.ToLower(strings.TrimSpace(response))
	return r == "yes" || r == "y"
}

func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
