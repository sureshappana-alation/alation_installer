package main

import (
	"fmt"
)

const ( // bash commands used in this file
	nodeNameCmd          = "kubectl get node --selector='node-role.kubernetes.io/master' -o jsonpath='{.items..metadata.name}'"
	applyStorageClassCmd = "kubectl apply -f res/kurl_patch/local-storage.yaml"

	labelMonitoringNodeCmd  = "kubectl label node %s node-role.alation.com/monitoring=monitoring --overwrite"
	prometheusDeletePvcCmd  = "kubectl -n monitoring delete pvc prometheus-k8s-db-prometheus-k8s-%s"
	prometheusMkdirCmd      = "sudo sudo mkdir -p /mnt/disks/prometheus-db-0"
	applyPrometheusPatchCmd = "kubectl apply -f res/kurl_patch/prometheus.yaml"

	labelRegistryNodeCmd  = "kubectl label node %s node-role.alation.com/registry=registry --overwrite"
	registryDeletePvcCmd  = "kubectl -n kurl delete pvc registry-pvc"
	registryMkdirCmd      = "sudo mkdir -p /mnt/disks/registry"
	applyRegistryPatchCmd = "kubectl apply -f res/kurl_patch/registry.yaml"
)

// The Kurl bootstrapper build does not include a storage solution and the cluster would need extra configuration for storage
// including creation of a storageClass and configuration of prometheus statefulset pv and pvc
func configClusterPlugins() {
	// Get node name
	_, nodeName := RunBashCmd(nodeNameCmd)
	LOGGER.Info("Cluster node name is " + nodeName)

	setupLocalStorageClass()

	setupPrometheus(nodeName)

	setupRegistry(nodeName)
}

func setupLocalStorageClass() {
	_, out := RunBashCmd(applyStorageClassCmd)
	LOGGER.Info(out)
}

// TODO - improve logic and logging
func setupPrometheus(nodeName string) {

	// Label node for monitoring
	_, out := RunBashCmd(fmt.Sprintf(labelMonitoringNodeCmd, nodeName))
	LOGGER.Info(out)

	// Delete existing persistent volume claims
	_, out = RunBashCmd(fmt.Sprintf(prometheusDeletePvcCmd, "0"))
	LOGGER.Info(out)
	_, out = RunBashCmd(fmt.Sprintf(prometheusDeletePvcCmd, "1"))
	LOGGER.Info(out)

	// TODO - config statfulset #replica

	// Create directory for prometheus db data
	_, out = RunBashCmd(prometheusMkdirCmd)
	LOGGER.Info(out)

	// Apply prometheus manifests
	_, out = RunBashCmd(applyPrometheusPatchCmd)
	LOGGER.Info(out)
}

func setupRegistry(nodeName string) {

	// Label node for registry
	_, out := RunBashCmd(fmt.Sprintf(labelRegistryNodeCmd, nodeName))
	LOGGER.Info(out)

	// Delete existing persistent volume claims
	_, out = RunBashCmd(registryDeletePvcCmd)
	LOGGER.Info(out)

	// Create directory for registry db data
	_, out = RunBashCmd(registryMkdirCmd)
	LOGGER.Info(out)

	// Apply registry manifests
	_, out = RunBashCmd(applyRegistryPatchCmd)
	LOGGER.Info(out)
}
