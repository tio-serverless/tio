package main

type B struct {
	Root      string
	BuildInfo buildInfo `toml:"build"`
}

type buildInfo struct {
	Name string `toml:"name"`
}
