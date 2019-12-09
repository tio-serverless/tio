package data

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"tio/database"
)

type B struct {
	Log         string  `toml:"log"`
	RestPort    int     `toml:"rest_port"`
	BuildAgent  string  `toml:"build_agent_address"`
	DeployAgent string  `toml:"deploy_agent_address"`
	RpcProt     int     `toml:"rpc_port"`
	Storage     storage `toml:"storage"`
	DBInfo      dbInfo  `toml:"db"`
	DBCli       database.TioDb
}

type dbInfo struct {
	Engine  string `toml:"engine"`
	Connect string `toml:"connect"`
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

	if b.DBInfo.Engine != "" {
		dc, err := database.GetDBClient(b.DBInfo.Engine, b.DBInfo.Connect)
		if err != nil {
			return nil, err
		}

		b.DBCli = dc
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
	logrus.Printf("RPC Port: %d", b.RpcProt)
	logrus.Printf("Build Agent Address: %s", b.BuildAgent)
	logrus.Printf("Deploy Agent Address: %s", b.DeployAgent)
	if b.DBCli != nil {
		logrus.Printf("DB Engine: %s", b.DBCli.Version())
	}
	logrus.Println("Storage:")
	logrus.Printf("  Acess Key: %s", b.Storage.AcessKey)
	logrus.Printf("  Sceret Key: %s", b.Storage.SecretKey)
	logrus.Printf("  Domain: %s", b.Storage.Domain)
	logrus.Println("----------------------")
}
