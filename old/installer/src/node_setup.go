package main

import "fmt"

const (
	registryConfigOnContainerdCmds = `
sudo sed -i 's|\[plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"\]|\[plugins."io.containerd.grpc.v1.cri".registry.mirrors."registry.kurl"\]|' /etc/containerd/config.toml;
sudo sed -i "s|https://registry-1.docker.io|http://$(kubectl -n kurl get service registry -o=jsonpath='{@.spec.clusterIP}')|" /etc/containerd/config.toml;
sudo systemctl stop containerd;
sudo systemctl start containerd;`

	nodeNameCmd  = "kubectl get node --selector='node-role.kubernetes.io/master' -o jsonpath='{.items..metadata.name}'"
	labelNodeCmd = "kubectl label node %s node-role.alation.com/%s --overwrite"
)

var nodeLabels = []string{"monitoring", "registry", "analytics", "catalog", "ocf"}

func setupNodes() {
	labelNodes()
	configRegistryOnContainerd()
}

// current setup is for single node cluster and all labels apply to the one node
func labelNodes() {
	// Get node name
	_, nodeName := RunBashCmd(nodeNameCmd)
	LOGGER.Info("Cluster node name is " + nodeName)

	for _, label := range nodeLabels {
		// Label node for monitoring
		_, out := RunBashCmd(fmt.Sprintf(labelNodeCmd, nodeName, fmt.Sprintf("%s=%s", label, "labeled")))
		LOGGER.Info(out)
	}

}

// This function config containerd to refer to the registry deployed on the cluster instead of docker.io registry
func configRegistryOnContainerd() {
	_, out := RunBashCmd(registryConfigOnContainerdCmds)
	LOGGER.Info(out)
}
