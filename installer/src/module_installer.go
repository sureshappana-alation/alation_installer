package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	persistentVolumeTemplate = "res/kurl_patch/persistent-volume-template.yaml"
	modulesDirPath           = "res/modules"
	moduleChartsPath         = "/helm.tgz"
	moduleImagesPath         = "/images"
)

const ( // Bash commands used in this file
	containerdLoadImageCmd = "sudo ctr -n k8s.io images import %s"
	helmInstallCmd         = "/usr/local/bin/helm upgrade --install %s %s %s"
	kubeApplyFileCmd       = "cat <<EOF | kubectl apply -f -\n%s\nEOF"
	mkdirCmd               = "sudo mkdir -p %s"
)

// This function looks at the modules available in res/module directory and install them all
func InstallModules(installConfig AlationInstallConfig) {

	var modulePaths []string

	// scan modules directory to find available modules
	moduleScanError :=
		filepath.WalkDir(modulesDirPath, func(path string, dirInfo os.DirEntry, err error) error {
			if dirInfo.IsDir() && path != modulesDirPath {
				modulePaths = append(modulePaths, path)
				return filepath.SkipDir // to only scan top level directories under modules directory
			}
			return nil
		})

	if moduleScanError != nil {
		logAndShowError("Error in scanning installer modules directory. %s", moduleScanError.Error())
		os.Exit(1)
	}

	// install modules
	// TODO - add modules installation priority logic
	for _, modulePath := range modulePaths {
		installModule(modulePath, installConfig)
	}
}

func installModule(modulePath string, installConfig AlationInstallConfig) {
	moduleName := filepath.Base(modulePath)
	logAndShowMsg(fmt.Sprintf("Module %s installation is started.", moduleName))

	loadModuleImages(modulePath)
	createPersistentVolumes(modulePath)
	installModuleCharts(modulePath, installConfig)
}

// Loads container images of the module into containerd
func loadModuleImages(modulePath string) {
	var imageTarBallPaths []string

	imagesDirPath := modulePath + moduleImagesPath

	// find all images of the module
	imagesDirScanError :=
		filepath.Walk(imagesDirPath, func(filePath string, fileInfo os.FileInfo, err error) error {
			if !fileInfo.IsDir() && filepath.Ext(filePath) == ".tar" {
				imageTarBallPaths = append(imageTarBallPaths, filePath)
			}
			return nil
		})

	if imagesDirScanError != nil {
		logAndShowError("Error in scanning modules images directory. %s", imagesDirScanError.Error())
		os.Exit(1)
	}

	// load images onto containerd
	for _, imageTarBallPath := range imageTarBallPaths {
		imageLoaded, out := RunBashCmd(fmt.Sprintf(containerdLoadImageCmd, imageTarBallPath))
		if imageLoaded {
			LOGGER.Info(fmt.Sprintf("Container image loaded: %s", imageTarBallPath))
		} else {
			logAndShowError("Module installer failed", fmt.Sprintf("Container image failed to load: %s", out))
			os.Exit(1)
		}
	}
}

// Install module helm charts
func installModuleCharts(modulePath string, installConfig AlationInstallConfig) {
	moduleName := filepath.Base(modulePath)

	valueOverrides := ""

	overriddenConf := installConfig.Modules[moduleName]
	for conf := range overriddenConf {
		value := installConfig.Modules[moduleName][conf]["value"]
		if value != "" {
			valueOverrides += fmt.Sprintf("--set %s=%s ", conf, value)
		}
	}

	installed, out := RunBashCmd(fmt.Sprintf(helmInstallCmd, valueOverrides, moduleName, modulePath+moduleChartsPath))

	if installed {
		logAndShowSuccess(fmt.Sprintf("Module %s is successfully installed.", moduleName))
	} else {
		logAndShowError(fmt.Sprintf("Module %s failed to install.", moduleName), "Error in applying charts.")
		LOGGER.Error(out) // log complete output
		os.Exit(1)
	}
}

// Create persistent volumes for the module base on the storage.yaml values.
func createPersistentVolumes(modulePath string) {
	moduleName := filepath.Base(modulePath)
	template, err := ioutil.ReadFile(persistentVolumeTemplate)
	if err != nil {
		logAndShowError("Module installation.", "Persistent volume template not found.")
		os.Exit(1)
	}

	storageFilePath := modulePath + "/storage.yaml"
	if fileExists(storageFilePath) {
		storage := ParseModuleStorage(storageFilePath)
		for _, volume := range storage.Volumes {

			// Create directory for volume
			_, out := RunBashCmd(fmt.Sprintf(mkdirCmd, volume.Path))
			LOGGER.Info(out)

			manifest := fmt.Sprintf(string(template), volume.Name, volume.Capacity, volume.Path, volume.Label, "labeled")

			persistentVolumeConfigured, out := RunBashCmd(fmt.Sprintf(kubeApplyFileCmd, manifest))
			if persistentVolumeConfigured {
				LOGGER.Info(fmt.Sprintf("Persistent volume %s configured for module %s: %s", volume.Name, moduleName, out))
			} else {
				logAndShowError("Module installer failed",
					fmt.Sprintf("Persistent volume %s not configured for module %s", volume.Name, moduleName))
				LOGGER.Error(out) // log complete output
				os.Exit(1)
			}
		}
	}
}
