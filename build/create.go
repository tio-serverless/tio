package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func buildImage(name string) error {
	return b.DClient.BuildImage(docker.BuildImageOptions{
		Name:       fmt.Sprintf("%s:%s", "vikings/tio-go-runtime", name),
		Dockerfile: b.Root + "/tio/Dockerfile",
		ContextDir: b.Root + "/tio/",
	})
}
