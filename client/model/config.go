package model

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	User User `toml:"user"`
	//Repo Repostry `toml:"repo"`
}

type User struct {
	Name   string `toml:"name"`
	Uid    int    `toml:"uid"`
	Passwd string `toml:"passwd"`
}

type Repostry struct {
	Url string `toml:"url"`
}

type MetaData struct {
	BuildInfo buildInfo `toml:"build"`
}

type buildInfo struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	API     string `toml:"api"`
	Rate    int32  `toml:"rate"`
	Stype   string `toml:"type"`
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

	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("$HOME/.tio Not Exist")
		os.MkdirAll(dir, 0700)
	}

	return ioutil.WriteFile(path, []byte(buf.String()), 0700)
}
