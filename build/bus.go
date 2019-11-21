package main

import docker "github.com/fsouza/go-dockerclient"

type B struct {
	Root      string
	Registry  string
	User      string
	Passwd    string
	BuildInfo buildInfo `toml:"build"`
	DClient   *docker.Client
}

type buildInfo struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

func dclientInit() error {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return err
	}

	b.DClient = client
	b.Registry = "vikings/tio-go-runtime"

	return nil
}
