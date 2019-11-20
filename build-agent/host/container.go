package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/sirupsen/logrus"
)

func runContainer(name, image string, cmd []string) error {
	c, err := b.DClient.CreateContainer(docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image: image,
			Cmd:   cmd,
		},
		HostConfig: &docker.HostConfig{
			Binds: []string{
				fmt.Sprintf("%s:/run/docker.sock", b.Build.Mount),
			},
		},
	})

	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	logrus.Debugf("Container ID: %s", c.ID)

	return b.DClient.StartContainer(c.ID, nil)
}
