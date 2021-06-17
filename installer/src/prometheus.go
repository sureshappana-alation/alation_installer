package main

import (
	"fmt"
	"os/exec"
)

func setupPrometheus(nodeName string) {

	// Label node for monitoring
	_, out := run_command(exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo kubectl label node %s node-role.alation.com/monitoring=monitoring --overwrite", nodeName)))
	LOGGER.Info(out)

	// Delete existing persistent volume claims
	_, out = run_command(exec.Command("/bin/sh", "-c", "sudo kubectl -n monitoring delete pvc prometheus-k8s-db-prometheus-k8s-0"))
	LOGGER.Info(out)
	_, out = run_command(exec.Command("/bin/sh", "-c", "sudo kubectl -n monitoring delete pvc prometheus-k8s-db-prometheus-k8s-1"))
	LOGGER.Info(out)

	// Create directory for prometheus db data
	_, out = run_command(exec.Command("/bin/sh", "-c", "sudo mkdir -p /mnt/disks/prometheus-db-0"))
	LOGGER.Info(out)

	// Apply prometheus manifests
	_, out = run_command(exec.Command("/bin/sh", "-c", "sudo kubectl apply -f res/manifests/prometheus.yaml"))
	LOGGER.Info(out)
}
