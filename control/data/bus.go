package data

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type B struct {
	Log        string  `toml:"log"`
	RestPort   int     `toml:"rest_port"`
	BuildAgent string  `toml:"build_agent_address"`
	Storage    storage `toml:"storage"`
}

type storage struct {
	AcessKey  string `toml:"accessKey"`
	SecretKey string `toml:"secretKey"`
	Domain    string `toml:"domain"`
}

func InitBus(file string) (*B, error) {
	b := new(B)
	_, err := toml.DecodeFile(file, b)
	if err != nil {
		return nil, err
	}

	if b.Storage.AcessKey == "" {
		b.Storage.AcessKey = os.Getenv("TIO_CONTROL_S_AKEY")
	}
	if b.Storage.SecretKey == "" {
		b.Storage.SecretKey = os.Getenv("TIO_CONTROL_S_SKEY")
	}
	if b.Storage.Domain == "" {
		b.Storage.Domain = os.Getenv("TIO_CONTROL_S_DOMAIN")
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
	logrus.Printf("Build Agent Address: %s", b.BuildAgent)
	logrus.Println("Storage:")
	logrus.Printf("  Acess Key: %s", b.Storage.AcessKey)
	logrus.Printf("  Sceret Key: %s", b.Storage.SecretKey)
	logrus.Printf("  Domain: %s", b.Storage.Domain)
	logrus.Println("----------------------")
}
