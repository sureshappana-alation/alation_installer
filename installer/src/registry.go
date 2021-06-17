package main

import (
	"fmt"
	"os/exec"
)

func setupRegistry(nodeName string) {

	// Label node for monitoring
	_, out := run_command(exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo kubectl label node %s node-role.alation.com/registry=registry --overwrite", nodeName)))
	LOGGER.Info(out)

	// Delete existing persistent volume claims
	_, out = run_command(exec.Command("/bin/sh", "-c", "sudo kubectl -n kurl delete pvc registry-pvc"))
	LOGGER.Info(out)

	// Create directory for local registry data
	_, out = run_command(exec.Command("/bin/sh", "-c", "sudo mkdir -p /mnt/disks/registry"))
	LOGGER.Info(out)

	// Apply registry manifests
	_, out = run_command(exec.Command("/bin/sh", "-c", "sudo kubectl apply -f res/manifests/registry.yaml"))
	LOGGER.Info(out)
}
