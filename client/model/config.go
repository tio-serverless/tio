package model

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	User User `toml:"user"`
	//Repo Repostry `toml:"repo"`
}

type User struct {
	Uid int `toml:"uid"`
}

type Repostry struct {
	Url string `toml:"url"`
}

func ReadConf(path string) (c Config, err error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Config{}, nil
	}

	var _c Config
	_, err = toml.DecodeFile(path, &_c)

	return _c, nil
}

func UpdateConf(c Config, path string) error {
	var buf bytes.Buffer

	e := toml.NewEncoder(&buf)

	err := e.Encode(c)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}

	return ioutil.WriteFile(path, []byte(buf.String()), 0700)
}
