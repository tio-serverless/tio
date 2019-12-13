package db

import (
	"errors"
	"strings"
	"time"

	"tio/control/data"
	"tio/database/model"
)

func SaveNewSrv(b *data.B, uid int, name string) (int, error) {
	s := model.Server{
		Name:      name,
		Uid:       uid,
		Status:    model.SrvBuilding,
		Domain:    "api.tio.io",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := b.DBCli.SaveTioServer(&s)
	if err != nil {
		return 0, err
	}

	ns, err := b.DBCli.QueryTioServerByName(name)
	if err != nil {
		return 0, err
	}

	return ns.Id, nil
}

func UpdateSrvBuildResult(b *data.B, sid, status int, name, path, image, raw, stype, version string) error {
	ns, err := b.DBCli.QueryTioServerById(sid)
	if err != nil {
		return err
	}

	if ns.Name == "" {
		return errors.New("Can not find this serivce record ")
	}

	ns.Status = status
	ns.Path = path
	ns.Image = image
	ns.Name = name
	ns.Raw = raw
	ns.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	switch strings.ToLower(stype) {
	case "http":
		ns.Stype = 1
	case "grpc":
		ns.Stype = 0
	default:
		ns.Stype = 2
	}
	ns.Version = version
	return b.DBCli.UpdateTioServer(ns)
}

func UpdateSrvStatus(b *data.B, sid, stauts int) error {
	ns, err := b.DBCli.QueryTioServerById(sid)
	if err != nil {
		return err
	}

	if ns.Name == "" {
		return errors.New("Can not find this serivce record ")
	}

	ns.Status = stauts
	ns.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	return b.DBCli.UpdateTioServer(ns)
}

func UpdateSrvImage(b *data.B, sid int, image string) error {
	ns, err := b.DBCli.QueryTioServerById(sid)
	if err != nil {
		return err
	}

	ns.Image = image

	return b.DBCli.UpdateTioServer(ns)
}

func QueryUserAllSrv(b *data.B, uid, limit int, name string) ([]model.Server, error) {
	return b.DBCli.QueryTioServerByUser(uid, limit, name)
}

func QuerySrvById(b *data.B, sid int) (*model.Server, error) {
	return b.DBCli.QueryTioServerById(sid)
}

func RemoveSrvByID(b *data.B, sid int) error {
	ns, err := b.DBCli.QueryTioServerById(sid)
	if err != nil {
		return err
	}

	return b.DBCli.DeleteTioServer(ns.Name)
}

func RemoveSrvByName(b *data.B, name string) error {
	return b.DBCli.DeleteTioServer(name)
}
