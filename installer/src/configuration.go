/**
* This file contains code related to loading, parsing and validating the initail configuration passed to the installer script.
 */
package main

import (
	"flag"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Structs to keep the configuration needed for installer
type AlationInstallConfig struct {
	// Cluster Cluster                      `yaml:"installer"` // Not used until multi-node cluster support
	Modules map[string]map[string]map[string]string `yaml:"modules"` // configurations related to each module
}

// Not used until multi-node cluster support
/* type Cluster struct {
	ControlPlaneHosts []Node `yaml:"controlPlaneHost"`
	DataPlaneHosts    []Node `yaml:"dataPlaneHosts"`
}

type Node struct {
	NodeName   string   `yaml:"nodeName"`
	SshUser    string   `yaml:"sshUser"`
	SshKeyPath string   `yaml:"sshKeyPath"`
	Labels     []string `yaml:"labels"`
} */

// Load alation install config yaml file and parse it as AlationInstallConfig struct
func ParseAlationInstallConfig(configFilePath string) AlationInstallConfig {

	yamlFile, error := ioutil.ReadFile(configFilePath)

	if error != nil {
		logAndShowError("Loading alation install config file", error.Error())
		os.Exit(1)
	}

	config := AlationInstallConfig{}

	error = yaml.Unmarshal(yamlFile, &config)
	if error != nil {
		logAndShowError("Parsing alation install config file", error.Error())
		os.Exit(1)
	}

	return config
}

func PrepareInstallConfig() AlationInstallConfig {

	// Parse and validate yaml configurations
	config := ParseAlationInstallConfig("res/alation-install.yaml")
	LOGGER.Info("Parsed Alation install config YAML file: ", config)

	type Secret struct {
		module   string
		conf     string
		required bool
		value    *string
	}

	var secretConfs []Secret
	var secretConfsWithValueRef []Secret

	// find configurations of type secret
	for module, moduleConf := range config.Modules {
		for conf, confAttr := range moduleConf {
			if confAttr["secret"] == "True" {
				secretConfs = append(secretConfs, Secret{module: module, conf: conf, required: confAttr["required"] == "True"})
			}
		}
	}

	// Receive secret arguments from command line arguments
	for _, secretConf := range secretConfs {
		secretConf.value = flag.String(secretConf.module+"."+secretConf.conf, "", "Absolute path of the YAML config file.")
		secretConfsWithValueRef = append(secretConfsWithValueRef, secretConf)
	}
	flag.Parse()

	// set secrets to the final config
	for _, secretConf := range secretConfsWithValueRef {
		if *secretConf.value != "" {
			config.Modules[secretConf.module][secretConf.conf]["value"] = *secretConf.value
		} else if secretConf.required {
			logAndShowError("Required argument configuration missing.", secretConf.conf+" is missing.")
			os.Exit(1)
		}
	}

	// verify all required argument are present
	for _, moduleConf := range config.Modules {
		for _, confAttr := range moduleConf {
			if confAttr["required"] == "True" && confAttr["value"] == "" {
				logAndShowError("Required yaml configuration.", "Yaml configuration is missing.")
				os.Exit(1)
			}
		}
	}

	logAndShowMsg("Alation install config file found and parsed.")
	return config
}

//
// modules storage configurations
//

// Structs to keep the configuration needed for modules storages
type ModuleStorage struct {
	Volumes []Volume `yaml:"volumes"`
}

type Volume struct {
	Name     string `yaml:"name"`
	Path     string `yaml:"path"`
	Capacity string `yaml:"capacity"`
	Label    string `yaml:"label"`
}

// Load module storgae yaml file and parse it as ModuleStorage struct
func ParseModuleStorage(configFilePath string) ModuleStorage {

	yamlFile, error := ioutil.ReadFile(configFilePath)

	if error != nil {
		logAndShowError("Loading module storage file", error.Error())
		os.Exit(1)
	}

	storage := ModuleStorage{}

	error = yaml.Unmarshal(yamlFile, &storage)
	if error != nil {
		logAndShowError("Parsing module storage file", error.Error())
		os.Exit(1)
	}

	return storage
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
