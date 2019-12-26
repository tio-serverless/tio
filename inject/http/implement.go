package main

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

type gInject struct {
	redCli *redis.Client
}

func (g *gInject) Store(name string, urls []string) error {

	data, _ := json.Marshal(urls)
	g.redCli.Set(name, data, 0)
	return nil
}

func NewInject(add, passwd string, db int) (*gInject, error) {

	gi := &gInject{}

	gi.redCli = redis.NewClient(&redis.Options{
		Addr:     add,
		Password: passwd,
		DB:       db,
	})

	_, err := gi.redCli.Ping().Result()
	if err != nil {
		return nil, err
	}

	return gi, nil
}
