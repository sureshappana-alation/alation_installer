package main

import (
	"fmt"
)

const ( // bash commands used in this file
	applyStorageClassCmd = "kubectl apply -f res/kurl_patch/local-storage.yaml"
	podStatusCmd         = "kubectl -n %s get pod %s  -o jsonpath='{.status.phase}'"

	prometheusDeletePvcCmd  = "kubectl -n monitoring delete pvc prometheus-k8s-db-prometheus-k8s-%s"
	prometheusMkdirCmd      = "sudo sudo mkdir -p /mnt/disks/prometheus-db-0"
	applyPrometheusPatchCmd = "kubectl apply -f res/kurl_patch/prometheus.yaml"

	registryPodNameCmd        = "kubectl -n kurl get pods -l app=registry -o jsonpath='{.items..metadata.name}'"
	registryDeletePvcCmd      = "kubectl -n kurl delete pvc registry-pvc"
	registryMkdirCmd          = "sudo mkdir -p /mnt/disks/registry"
	applyRegistryPatchCmd     = "kubectl -n kurl apply -f res/kurl_patch/registry.yaml"
	setRegistryServicePortCmd = `kubectl -n kurl patch svc registry --type merge -p '{"spec":{"ports": [{"port": 80,"name":"registry","targetPort":80}]}}'`
)

// The Kurl bootstrapper build does not include a storage solution and the cluster would need extra configuration for storage
// including creation of a storageClass and configuration of prometheus statefulset pv and pvc
func configClusterPlugins() {

	setupLocalStorageClass()

	setupPrometheus()

	setupRegistry()
}

func setupLocalStorageClass() {
	_, out := RunBashCmd(applyStorageClassCmd)
	LOGGER.Info(out)
}

// TODO - improve logic and logging
func setupPrometheus() {

	_, podStatus := RunBashCmd(fmt.Sprintf(podStatusCmd, "monitoring", "prometheus-k8s-0"))

	if podStatus == "Pending" {
		// Delete existing persistent volume claims
		_, out := RunBashCmd(fmt.Sprintf(prometheusDeletePvcCmd, "0"))
		LOGGER.Info(out)
		_, out = RunBashCmd(fmt.Sprintf(prometheusDeletePvcCmd, "1"))
		LOGGER.Info(out)

		// TODO - config statefulset #replica

		// Create directory for prometheus db data
		_, out = RunBashCmd(prometheusMkdirCmd)
		LOGGER.Info(out)

		// Apply prometheus manifests
		_, out = RunBashCmd(applyPrometheusPatchCmd)
		LOGGER.Info(out)
	}
}

func setupRegistry() {

	_, podStatus := RunBashCmd(fmt.Sprintf(podStatusCmd, "kurl", fmt.Sprintf("$(%s)", registryPodNameCmd)))

	if podStatus == "Pending" {
		// Delete existing persistent volume claims
		_, out := RunBashCmd(registryDeletePvcCmd)
		LOGGER.Info(out)

		// Create directory for registry db data
		_, out = RunBashCmd(registryMkdirCmd)
		LOGGER.Info(out)

		// Apply registry manifests
		_, out = RunBashCmd(applyRegistryPatchCmd)
		LOGGER.Info(out)

		// Set registry service ports
		_, out = RunBashCmd(setRegistryServicePortCmd)
		LOGGER.Info(out)
	}
}
