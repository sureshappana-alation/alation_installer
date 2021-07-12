package main

const (
	registryConfigOnContainerdCmds = `
sudo sed -i 's|\[plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"\]|\[plugins."io.containerd.grpc.v1.cri".registry.mirrors."registry.kurl"\]|' /etc/containerd/config.toml;
sudo sed -i "s|https://registry-1.docker.io|http://$(kubectl -n kurl get service registry -o=jsonpath='{@.spec.clusterIP}')|" /etc/containerd/config.toml;
sudo systemctl stop containerd;
sudo systemctl start containerd;`
)

func configRegistryOnContainerd() {
	_, out := RunBashCmd(registryConfigOnContainerdCmds)
	LOGGER.Info(out)
}
