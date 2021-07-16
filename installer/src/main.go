package main

import (
	_ "reflect"
	_ "runtime"
)

func main() {
	logAndShowMsg("Installer Script started.")

	// TODO - Add environment verifications

	installConfig := PrepareInstallConfig()

	BootstrapKubernetesCluster()

	setupNodes()

	configClusterPlugins()

	InstallModules(installConfig)

	logAndShowSuccess("Installer Script finished successfully. \nInstaller logs are available at " + getLogFilePath())
}
