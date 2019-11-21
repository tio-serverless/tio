package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

func buildImage(name string) error {
	var buf bytes.Buffer

	f, err := os.Open(b.Root + "/tio.tar")
	if err != nil {
		log.Fatal(err)
	}

	return b.DClient.BuildImage(docker.BuildImageOptions{
		Name:                fmt.Sprintf("%s:%s", b.Registry, name),
		Dockerfile:          "Dockerfile",
		RmTmpContainer:      true,
		ForceRmTmpContainer: true,
		OutputStream:        &buf,
		InputStream:         f,
	})
}
