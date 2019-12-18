package model

const (
	SrvReady = iota
	SrvBuildSuc
	SrvBuildFailed
	SrvBuilding
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
	Timestamp string
	Status    int
	Image     string
	Raw       string
}

type User struct {
	Id     int
	Name   string
	Passwd string
}
