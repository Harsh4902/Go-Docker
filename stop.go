package main

import (
	"context"
	"fmt"

	containertypes "github.com/docker/docker/api/types/container"
)

func StopDocker(containerId string) {
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

	fmt.Print("Stopping container ", containerId, "... ")
	noWaitTimeout := 0 // to not wait for the container to exit gracefully
	if err := cli.ContainerStop(ctx, containerId, containertypes.StopOptions{Timeout: &noWaitTimeout}); err != nil {
		panic(err)
	}
	fmt.Println("Success")
}
