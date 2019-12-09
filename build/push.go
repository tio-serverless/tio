package main

import docker "github.com/fsouza/go-dockerclient"

func push(tag string) error {
	return b.DClient.PushImage(docker.PushImageOptions{
		Name: b.Registry,
		Tag:  tag,
	}, docker.AuthConfiguration{
		Username: b.User,
		Password: b.Passwd,
	})
}
