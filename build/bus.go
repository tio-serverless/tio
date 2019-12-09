package main

import docker "github.com/fsouza/go-dockerclient"

type B struct {
	Root      string
	Registry  string
	User      string
	Passwd    string
	J         *job
	UserName  string    `toml:"user"`
	BuildInfo buildInfo `toml:"build"`
	DClient   *docker.Client
}

type buildInfo struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	API     string `toml:"api"`
	Rate    int32  `toml:"rate"`
}

type job struct {
	User   string
	Name   string
	Image  string
	API    string
	Rate   int32
	Status int
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
