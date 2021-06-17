package main

import "os"

const (
	ENV_PATH     = "/opt/replicated/kubeadm.conf"
	ROOT_USER_ID = 0
)

func verify_root_access() {
	if os.Geteuid() != ROOT_USER_ID {
		show_error("User role check", "No root access. Run with root user or use sudo.")
		panic("no root access")
	}
}

func verify_first_installation() {
	check_file, err := os.Open(ENV_PATH)
	if err == nil {
		show_error("Kubernetes fresh installation check", "Kubernetes has previously installed on this machine.")
		panic("fail to install Kubernetes")
	}
	defer check_file.Close()
}
