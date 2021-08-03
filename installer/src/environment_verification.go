package main

import (
	"os"
	"os/exec"
	"strings"
)

func verifyEnvironment() {
	verifySudoAccess()
}

func verifySudoAccess() {
	// to check the root access will run a sudo command
	access, out := RunCommand(exec.Command("sudo", "echo", "sudoaccesschecked"), "Sudo access check.")

	if !access || !strings.Contains(out, "sudoaccesschecked") {
		logAndShowError("User access check", "Run with a user that has sudo access.")
		os.Exit(1)
	}
}
