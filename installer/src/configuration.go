/**
* This file contains code related to loading, parsing and validating the initail configuration passed to the installer script.
 */
package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Host labels used for node affinity for the deployment of services
var ValidHostLabels = []string{
	"monitoring",
	"registry",
	"analytics",
	"rosemeta",
	"rosemeta-replica",
}

// Structure to keep the configuration
type InstallerConfig struct {
	Version           string   `yaml:"version"`
	ControlPlaneHosts []Host   `yaml:"controlPlaneHost"`
	DataPlaneHosts    []Host   `yaml:"dataPlaneHosts"`
	Rosemeta          Rosemeta `yaml:"rosemeta"`
}

type Host struct {
	HostName       string   `yaml:"hostName"`
	SshUser        string   `yaml:"sshUser"`
	SshKeyPath     string   `yaml:"sshKeyPath"`
	EbsVolumePaths []string `yaml:"ebsVolumePaths"`
	Labels         []string `yaml:"labels"`
}

type Rosemeta struct {
	DataPath     string `yaml:"dataPath"`
	ReplicaCount int
}

// Load configuration yaml file and parse it as InstallerConfig struct
func ParseYamlConfiguration(filePath string) InstallerConfig {

	yamlFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		show_error("Loading Configuration File", err.Error())
		panic(err)
	}

	config := InstallerConfig{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		show_error("Parsing Configuration File", err.Error())
		panic(err)
	}
	validationError := validate(config)
	if validationError != nil {
		show_error("Validating Configurations", validationError.Error())
		panic(err)
	}

	return config
}

// Validate the configuration values
func validate(installerConfig InstallerConfig) error {
	// TODO - complete this function
	// check labels
	hosts := append(installerConfig.ControlPlaneHosts, installerConfig.DataPlaneHosts...)
	for _, host := range hosts {
		for _, label := range host.Labels {
			var valid bool = false
			for _, validLabel := range ValidHostLabels {
				if validLabel == label {
					valid = true
				}
			}
			if !valid {
				return fmt.Errorf("invalid Host Label: %s. Valid values: %v", label, ValidHostLabels)
			}
		}
	}

	return nil
}
