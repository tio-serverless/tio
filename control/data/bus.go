package data

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type B struct {
	Log      string `toml:"log"`
	RestPort int    `toml:"rest_port"`
}

func InitBus(file string) (*B, error) {
	b := new(B)
	_, err := toml.DecodeFile(file, b)
	if err != nil {
		return nil, err
	}

	output(b)
	enableLog(b)

	return b, nil
}

func enableLog(b *B) {
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

func output(b *B) {
	logrus.Println("----------------------")
	logrus.Printf("Control Log: %s", b.Log)
	logrus.Printf("Rest Port: %d", b.RestPort)
	logrus.Println("----------------------")
}
