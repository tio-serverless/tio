package main

import docker "github.com/fsouza/go-dockerclient"

func runContainer(name, image string, cmd []string) error {
	c, err := b.DClient.CreateContainer(docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image: image,
			Cmd:   cmd,
		},
	})

	if err != nil {
		return err
	}

	return b.DClient.StartContainer(c.ID, nil)
}
