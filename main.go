package main

import (
	"os"
)

func main() {
	cmd := os.Args[1]

	if cmd == "start" {
		imageName := os.Args[2]
		StartDocker(imageName)
	}

	if cmd == "stop" {
		containerId := os.Args[2]
		StopDocker(containerId)
	}
}
