package main

import (
	"fmt"
	"os/exec"
)

// Run a bash command on host OS
// Each command will be run as a child process to the installer process
// Each command process will have a separate session with same environment variables of the installer process
// Environment changes of each command will not affect the main process or next commands call context
func RunBashCmd(cmd string) (bool, string) {
	return RunCommand(exec.Command("bash", "-c", cmd))
}

func RunCommand(cmd *exec.Cmd) (bool, string) {
					LOGGER.Info(fmt.Sprintf("running cmd: %s", cmd))
	out, err := cmd.CombinedOutput()
	if err != nil {
		errStr := fmt.Sprintf("%s%s", out, err)
		LOGGER.Error(fmt.Sprintf("Error in running command: %s", cmd))
		return false, errStr
	}
	return true, fmt.Sprintf("%s", out)
}
