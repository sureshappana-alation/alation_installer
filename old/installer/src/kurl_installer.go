package main

import (
	"fmt"
	"os"
	"strings"
)

const ( // Bash commands used in this file
	kubeClusterInfoCmd = "kubectl cluster-info"
	workingDirCmd      = "pwd"
	deleteDirCmd       = "if [ -d \"%s\" ]; then rm -Rf %s; fi"
	makeKurlDirCmd     = "mkdir %s/res/kurl"
	kurlDecompressCmd  = "tar xvzf %s/res/kurl.tar.gz -C %s/res/kurl"
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
}

func decompressKurlPackage() {
	// get working directory
	_, cmdOut := RunBashCmd(workingDirCmd)
	workingDirectory = strings.TrimSpace(cmdOut)
	LOGGER.Info("Working directory is", workingDirectory)

	kurlDirectory = fmt.Sprintf("%s/res/kurl", workingDirectory)
	LOGGER.Info("Desired path for decompressed kurl package is ", workingDirectory)

	// delete decompress kurl directory if exist - in case of installer re-run
	RunBashCmd(fmt.Sprintf(deleteDirCmd, kurlDirectory, kurlDirectory))

	// create directory to decompress kURL package into
	kurlDirectoryCreated, cmdOut := RunBashCmd(fmt.Sprintf(makeKurlDirCmd, workingDirectory))
	if !kurlDirectoryCreated {
		logAndShowError("Could not create new directory for Kubernetes bootstrapper.", cmdOut)
		os.Exit(1)
	}

	// Decompress kURL package
	decompressed, cmdOut :=
		RunBashCmd(fmt.Sprintf(kurlDecompressCmd, workingDirectory, workingDirectory))
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

	LOGGER.Info("Kurl install out: \n" + cmdOut)

	kubeConfig, kubeConfigExist := os.LookupEnv("KUBECONFIG")
	if kubeConfigExist {
		LOGGER.Info("KUBECONFIG environmnet already set. Value:" + kubeConfig)
	}
	// This environment variable is needed for next commands which work with kubectl
	os.Setenv("KUBECONFIG", "/etc/kubernetes/admin.conf")
	LOGGER.Info("KUBECONFIG environmnet variable set to /etc/kubernetes/admin.conf")

	logAndShowSuccess("Kubernetes Cluster installed and is running.")
}
