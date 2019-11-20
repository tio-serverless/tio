package main

import (
	"bytes"
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func buildImage(name string) error {
	var buf bytes.Buffer

	return b.DClient.BuildImage(docker.BuildImageOptions{
		Name:                fmt.Sprintf("%s:%s", "vikings/tio-go-runtime", name),
		Dockerfile:          b.Root + "/tio/Dockerfile",
		RmTmpContainer:      true,
		ForceRmTmpContainer: true,
		ContextDir:          b.Root + "/tio/",
		OutputStream:        &buf,
	})
}
