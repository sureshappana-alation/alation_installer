/**
* This file contains code related to loading, parsing and validating the initial configuration passed to the installer script.
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

// Structs to keep the configuration needed for installer
type AlationInstallConfig struct {
	// Cluster Cluster                      `yaml:"installer"` // Not used until multi-node cluster support
	Modules map[string]map[string]interface{} `yaml:"modules"` // configurations related to each module
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
		logAndShowError("Loading Alation install config file", error.Error())
		os.Exit(1)
	}

	config := AlationInstallConfig{}

	error = yaml.Unmarshal(yamlFile, &config)
	if error != nil {
		logAndShowError("Parsing Alation install config file", error.Error())
		os.Exit(1)
	}

	// verify config types
	// all module configuration should be type of key-values except for volumes which is an array of key-values
	confTypeError := false
	for moduleName, moduleConfs := range config.Modules {
		for confName, conf := range moduleConfs {
			switch conf.(type) {
			case map[string]interface{}:
				if confName == "volumes" {
					confTypeError = true
				}
			case []interface{}:
				if confName != "volumes" {
					confTypeError = true
				}
			default:
				confTypeError = true
			}
			if confTypeError {
				logAndShowError("alation-install.yaml configuration file",
					fmt.Sprintf("unknown configuration type. module=%s, property=%s, type=%s",
						moduleName, confName, reflect.TypeOf(conf)))
				os.Exit(1)
			}
		}
	}

	return config
}

func PrepareInstallConfig() AlationInstallConfig {

	// Parse and validate yaml configurations
	config := ParseAlationInstallConfig("alation-install.yaml")
	LOGGER.Info("Parsed Alation install config YAML file: ", config)

	type Secret struct {
		description string
		module      string
		conf        string
		required    bool
		value       string
	}

	var secretConfs []Secret
	var secretConfsWithValueRef []Secret

	// find configurations of type secret
	for module, moduleConf := range config.Modules {
		for confName, conf := range moduleConf {
			if confName != "volumes" {
				confMap := conf.(map[string]interface{})
				if confMap["secret"] == true {
					secretConfs = append(secretConfs,
						Secret{
							description: confMap["description"].(string),
							module:      module,
							conf:        confName,
							required:    confMap["required"] == true})
				}
			}
		}
	}

	// Receive secret arguments from command line arguments
	for _, secretConf := range secretConfs {
		secretConf.value =
			getFromEnvOrPromptSecret(secretConf.module+"."+secretConf.conf, secretConf.description, secretConf.required)
		secretConfsWithValueRef = append(secretConfsWithValueRef, secretConf)
	}

	// set secrets to the final config
	for _, secretConf := range secretConfsWithValueRef {
		if secretConf.value != "" {
			config.Modules[secretConf.module][secretConf.conf].(map[string]interface{})["value"] = secretConf.value
		} else if secretConf.required {
			logAndShowError("Required argument configuration missing.",
				fmt.Sprintf("modules.%s.%s is missing.", secretConf.module, secretConf.conf))
			os.Exit(1)
		}
	}

	// verify all required argument are present
	for moduleName, moduleConf := range config.Modules {
		for confName, conf := range moduleConf {
			if confName != "volumes" {
				confMap := conf.(map[string]interface{})
				if confMap["required"] == true && (confMap["value"] == nil || confMap["value"] == "") {
					logAndShowError("Required yaml configuration.",
						fmt.Sprintf("Yaml configuration modules.%s.%s value is missing.", moduleName, confName))
					os.Exit(1)
				}
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

// Load module storage yaml file and parse it as ModuleStorage struct
// volume paths get overridden from install config
func ParseModuleStorage(moduleName string, configFilePath string, installConfig AlationInstallConfig) ModuleStorage {

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
	overrideVolumes := installConfig.Modules[moduleName]["volumes"]

	if overrideVolumes != nil {
		overriddenVolumes := make([]Volume, 0)

		for _, volume := range storage.Volumes {
			var overriddenPath string = volume.Path
			for _, volumeOverride := range overrideVolumes.([]interface{}) {
				name := volumeOverride.(map[string]interface{})["name"]
				path := volumeOverride.(map[string]interface{})["path"]
				if volume.Name == name && volume.Path != path {
					LOGGER.Info(fmt.Sprintf("Volume path overriden with install configuration. module=%s, persistantVolume=%s, defaultPath=%s, newPath=%s",
						moduleName, volume.Name, volume.Path, path))
					overriddenPath = path.(string)
				}
			}
			overriddenVolumes = append(overriddenVolumes,
				Volume{Name: volume.Name, Capacity: volume.Capacity, Label: volume.Label, Path: overriddenPath})
		}

		return ModuleStorage{Volumes: overriddenVolumes}
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
