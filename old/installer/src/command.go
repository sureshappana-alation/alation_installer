package main

import (
	"fmt"
	"os/exec"
)

// Run a bash command on host OS
// Each command will be run as a child process to the installer process
// Each command process will have a separate session with same environment variables of the installer process
// Environment changes of each command will not affect the main process or next commands call context
// Command string will be logged as it is
func RunBashCmd(cmd string) (bool, string) {
	return RunCommand(exec.Command("bash", "-c", cmd), "")
}

// Run an exec.Cmd and log the message
// If logMsg is empty then command string will be logged
func RunCommand(cmd *exec.Cmd, logMsg string) (bool, string) {
	log := fmt.Sprintf("Running command: %s", cmd)
	if logMsg != "" {
		log = fmt.Sprintf("Running command with message: %s", logMsg)
	}

	LOGGER.Info(log)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errStr := fmt.Sprintf("%s - %s", out, err.Error())
		LOGGER.Error(fmt.Sprintf("Error in running command: %s", errStr))
		return false, errStr
	}
	return true, string(out)
}
