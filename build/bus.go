package main

import docker "github.com/fsouza/go-dockerclient"

type B struct {
	Root      string
	BuildInfo buildInfo `toml:"build"`
	DClient *docker.Client
}

type buildInfo struct {
	Name string `toml:"name"`
}

func dclientInit() error {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return err
	}

	b.DClient = client

	return nil
}

