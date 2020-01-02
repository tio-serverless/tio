package main

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

type meta struct {
	Url       string `json:"url"`
	RouteType int    `json:"route_type"`
	Remove    bool   `json:"remove"`
}

func (m monImplement) DisableService(name string) error {
	path := fmt.Sprintf("tio/v1/gateway/services/%s", name)
	logrus.Debugf("Disable %s ", path)
	val, _, err := m.consulCli.KV().Get(path, nil)
	if err != nil {
		return err
	}

	if val == nil {
		return nil
	}

	var mta meta
	json.Unmarshal(val.Value, &m)
	fmt.Println(m)

	mta.Remove = true

	content, _ := json.Marshal(m)

	_, err = m.consulCli.KV().Put(&api.KVPair{
		Key:   val.Key,
		Value: content,
	}, nil)

	return err
}
