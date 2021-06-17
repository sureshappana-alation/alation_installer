package main

import (
	"flag"
	"os/exec"
	_ "reflect"
	_ "runtime"
)

func main() {
	LOGGER.Info("Installer Script started.")

	// Initial environment checks
	verify_root_access()

	// Parse command-line arguments
	config_file_path := flag.String("config_file_path", "res/conf.yaml", "Absolute path of the YAML config file.")
	flag.Parse()
	LOGGER.Info("Parsed argument for config file Path: ", *config_file_path)

	// Parse and validate yaml configurations
	var config = ParseYamlConfiguration(*config_file_path)
	LOGGER.Info("Parsed config YAML file: ", config)
	show_result("Configuration file found and parsed successfully.")

	kubernetesInstalled, _ := run_command(exec.Command("/bin/sh", "-c", "sudo kubectl get nodes"))

	if !kubernetesInstalled { // improve this logic to check version
		// install Kubernetes cluster
		decompressKurlPackage()
		installKubernetes()
	} else {
		show_result("Kubernetes platform is already installed. Setup will be continued.")
	}

	// Get node name
	_, nodeName := run_command(exec.Command("/bin/sh", "-c", "sudo kubectl get node --selector='node-role.kubernetes.io/master' -o jsonpath='{.items..metadata.name}'"))
	LOGGER.Info("Cluster node name is " + nodeName)

	setupLocalStorageClass()
	setupPrometheus(nodeName)
	setupRegistry(nodeName)

	LOGGER.Info("Installer Script finished.")
}
