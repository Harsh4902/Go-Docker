package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

func GetClient(driver string) (*client.Client, error) {
	if driver != "docker" {
		out, err := exec.Command("podman", "machine", "inspect", "--format", "{{.ConnectionInfo.PodmanSocket.Path}}").Output()
		if err != nil {
			fmt.Println(err)
		}

		err = os.Setenv("DOCKER_HOST", "unix://"+strings.TrimSpace(string(out)))
		if err != nil {
			fmt.Println(err)
		}
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		return nil, err
	}

	return cli, nil
}
