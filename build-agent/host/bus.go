package main

import (
	"os"

	"github.com/BurntSushi/toml"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/sirupsen/logrus"
)

var b *bus

type bus struct {
	Log     string `toml:"log"`
	Port    int    `toml:"port"`
	Build   build  `toml:"build"`
	DClient *docker.Client
}

type build struct {
	Image string `toml:"build_image"`
	Base  string `toml:"base_image"`
	Mount string `toml:"mount"`
}

func init() {
	b = new(bus)
	_, err := toml.DecodeFile(os.Getenv("TIO_BUILD_CONFIG"), b)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	err = dclientInit()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	enableLog()

	output()

	return
}

func dclientInit() error {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return err
	}

	b.DClient = client

	return nil
}

func enableLog() {
	logrus.SetLevel(logrus.InfoLevel)
	switch b.Log {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}
}

func output() {
	i, _ := b.DClient.Info()
	logrus.Println("----------------------")
	logrus.Printf("Control Log: %s", b.Log)
	logrus.Printf("GRPC Port: %d", b.Port)
	logrus.Printf("Docker Client Version: %s", i.KernelVersion)
	logrus.Print("Build")
	logrus.Printf("  Build Image: %s", b.Build.Image)
	logrus.Printf("  Base Image: %s", b.Build.Base)
	logrus.Println("----------------------")
}
