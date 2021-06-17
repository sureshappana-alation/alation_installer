package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	KURL_BUILD_HASH = "d2b213e"
)

var (
	workingDirectory string
	kurlDirectory    string
)

func decompressKurlPackage() {
	// get working directory
	_, cmdOut := run_command(exec.Command("pwd"))
	workingDirectory = strings.TrimSpace(cmdOut)
	LOGGER.Info("Working directory is", workingDirectory)

	kurlDirectory = fmt.Sprintf("%s/res/kurl-%s", workingDirectory, KURL_BUILD_HASH)
	LOGGER.Info("Desired path for decompressed kurl package is ", workingDirectory)

	// delete the kurl decompress directory if exist - in case of installer re-run
	run_command(exec.Command("/bin/sh", "-c", fmt.Sprintf("if [ -d \"%s\" ]; then rm -Rf %s; fi", kurlDirectory, kurlDirectory)))

	// create directory to decompress kURL package into
	kurlDirectoryCreated, cmdOut := run_command(exec.Command("/bin/sh", "-c", fmt.Sprintf("mkdir %s/res/kurl-%s", workingDirectory, KURL_BUILD_HASH)))
	if !kurlDirectoryCreated {
		LOGGER.Error(cmdOut)
		panic(fmt.Sprintf("Could not create new directory for Kubernetes bootstraper, Error: %s", cmdOut))
	}
	LOGGER.Info(cmdOut)

	// Decompress kURL package
	decompressed, cmdOut := run_command(exec.Command("/bin/sh", "-c", fmt.Sprintf("tar xvzf %s/res/kurl-%s.tar.gz -C %s/res/kurl-%s", workingDirectory, KURL_BUILD_HASH, workingDirectory, KURL_BUILD_HASH)))
	if !decompressed {
		LOGGER.Error(cmdOut)
		panic(fmt.Sprintf("Could not decompress Kubernetes bootstraper package, Error: %s", cmdOut))
	}
	LOGGER.Info(cmdOut)
	show_result("Kubernetes Bootstrapper package decompressed.")
}

func installKubernetes() {
	kurlInstallerSucceed, cmdOut := run_command(exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s && cat install.sh | sudo bash -s airgap", kurlDirectory)))

	if !kurlInstallerSucceed {
		LOGGER.Error(cmdOut)
		show_error("Kubernetes platform installation", "failed")
		panic("Kubernetes platform installation failed")
	}
	LOGGER.Info(cmdOut)
	show_result("Kubernetes Cluster installed.")
}

func setupLocalStorageClass() {
	_, out := run_command(exec.Command("/bin/sh", "-c", "sudo kubectl apply -f res/manifests/local-storage.yaml"))
	LOGGER.Info(out)
}

func run_command(cmd *exec.Cmd) (bool, string) {
	LOGGER.Info(fmt.Sprintf("running cmd: %s", cmd))
	out, err := cmd.CombinedOutput()
	if err != nil {
		err_str := fmt.Sprintf("%s%s", out, err)
		LOGGER.Error(fmt.Sprintf("running cmd: %s", cmd))
		return false, err_str
	}
	return true, fmt.Sprintf("%s", out)
}
