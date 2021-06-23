package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	// spec of the build is available at https://kurl.sh/d2b213e
	// and https://alationcorp.atlassian.net/wiki/spaces/ENG/pages/5595541243/Alation+kURL+build
	kurlBuildHash = "d2b213e"
)

const ( // Bash commands used in this file
	kubeClusterInfoCmd = "kubectl cluster-info"
	workingDirCmd      = "pwd"
	deleteDirCmd       = "if [ -d \"%s\" ]; then rm -Rf %s; fi"
	makeKurlDirCmd     = "mkdir %s/res/kurl-%s"
	kurlDecompressCmd  = "tar xvzf %s/res/kurl-%s.tar.gz -C %s/res/kurl-%s"
	kurlInstallCmd     = "cd %s && cat install.sh | sudo bash -s airgap"
)

var (
	workingDirectory string
	kurlDirectory    string
)

// This function bootstraps Kubernetes cluster using kURL.sh air gap solution if not already installed/configured.
// Current implementation support only single node clusters.
func BootstrapKubernetesCluster() {
	kubernetesInstalled, _ := RunBashCmd(kubeClusterInfoCmd) // TODO - the check logic could improve

	if !kubernetesInstalled { // TODO - improve this logic to check version
		decompressKurlPackage()

		logAndShowMsg("Kubernetes installation started. This could take several minutes...")
		bootstrapKubernetes()

	} else {
		logAndShowMsg("Kubernetes platform is already installed. Alation Setup will be continued.")
	}

	// extra config needed to run cluster plugins after kurl installation
	configClusterPlugins()
}

func decompressKurlPackage() {
	// get working directory
	_, cmdOut := RunBashCmd(workingDirCmd)
	workingDirectory = strings.TrimSpace(cmdOut)
	LOGGER.Info("Working directory is", workingDirectory)

	kurlDirectory = fmt.Sprintf("%s/res/kurl-%s", workingDirectory, kurlBuildHash)
	LOGGER.Info("Desired path for decompressed kurl package is ", workingDirectory)

	// delete decompress kurl directory if exist - in case of installer re-run
	RunBashCmd(fmt.Sprintf(deleteDirCmd, kurlDirectory, kurlDirectory))

	// create directory to decompress kURL package into
	kurlDirectoryCreated, cmdOut := RunBashCmd(fmt.Sprintf(makeKurlDirCmd, workingDirectory, kurlBuildHash))
	if !kurlDirectoryCreated {
		logAndShowError("Could not create new directory for Kubernetes bootstrapper.", cmdOut)
		os.Exit(1)
	}

	// Decompress kURL package
	decompressed, cmdOut :=
		RunBashCmd(fmt.Sprintf(kurlDecompressCmd, workingDirectory, kurlBuildHash, workingDirectory, kurlBuildHash))
	if !decompressed {
		logAndShowError("Could not decompress Kubernetes bootstrapper package.", cmdOut)
		os.Exit(1)
	}
	logAndShowMsg("Kubernetes Bootstrapper package decompressed.")
}

func bootstrapKubernetes() {

	kurlInstallerSucceed, cmdOut := RunBashCmd(fmt.Sprintf(kurlInstallCmd, kurlDirectory))

	if !kurlInstallerSucceed {
		LOGGER.Error("Kurl install out: \n" + cmdOut)
		logAndShowError("Kubernetes platform installation", "failed")
		os.Exit(1)
	}

	// This environment variable is needed for next commands which work with kubectl
	os.Setenv("KUBECONFIG", "/etc/kubernetes/admin.conf")

	LOGGER.Info("Kurl install out: \n" + cmdOut)
	logAndShowSuccess("Kubernetes Cluster installed.")
}
