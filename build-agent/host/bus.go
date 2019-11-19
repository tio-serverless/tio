package main

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var b *bus

type bus struct {
	Log   string `toml:"log"`
	Port  int    `toml:"port"`
	Build build  `toml:"build"`
}

type build struct {
	Base string `toml:"base_image"`
}

func init() {
	b = new(bus)
	_, err := toml.DecodeFile(os.Getenv("TIO_BUILD_CONFIG"), b)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	enableLog()

	output()

	return
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
	logrus.Println("----------------------")
	logrus.Printf("Control Log: %s", b.Log)
	logrus.Printf("GRPC Port: %d", b.Port)
	logrus.Print("Build")
	logrus.Printf("  Base Image: %s", b.Build.Base)
	logrus.Println("----------------------")
}
