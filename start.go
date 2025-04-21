package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/go-connections/nat"
)

func StartDocker(imageName string) {
	ctx := context.Background()

	cfg, err := LoadConfig(ConfigPath)

	if err != nil {
		fmt.Println(err)
	}

	cli, err := GetClient(cfg.Instance.Driver)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Define exposed port and bindings
	exposedPort, _ := nat.NewPort("tcp", "80")
	portBindings := nat.PortMap{
		exposedPort: []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: "8080",
			},
		},
	}

	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image:        imageName,
			ExposedPorts: nat.PortSet{exposedPort: struct{}{}},
		},
		&container.HostConfig{
			PortBindings: portBindings,
		}, nil, nil, "")

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}
