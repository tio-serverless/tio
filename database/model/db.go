package model

import "time"

const (
	SrvBuilding = iota
	SrvBuildSuc
	SrvBuildFailed
	SrvDeploying
	SrvDeploySuc
	SrvDeployFailed
)

type Server struct {
	Id        int
	Name      string
	Version   string
	Uid       int
	Stype     int
	Domain    string
	Path      string
	TVersion  string
	Timestamp time.Time
	Status    int
	Raw       string
}

type User struct {
	Id     int
	Name   string
	Passwd string
}
