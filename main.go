package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := os.Args[1]

	if cmd == "start" {
		imageName := os.Args[2]
		driver := os.Args[3]
		cfg, err := LoadConfig(ConfigPath)

		if err != nil {
			fmt.Println(err)
		}

		cfg.Instance.Driver = driver
		SaveConfig(ConfigPath, cfg)
		StartDocker(imageName)
	}

	if cmd == "stop" {
		containerId := os.Args[2]
		StopDocker(containerId)
	}
}
