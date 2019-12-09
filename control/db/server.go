package db

import (
	"tio/control/data"
	"tio/database/model"
)

func SaveNewSrv(b *data.B, uid int, name string) (int, error) {
	s := model.Server{
		Name:   name,
		Uid:    uid,
		Status: model.SrvBuilding,
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

func UpdateSrvStatus(b *data.B, sid, stauts int) error {
	ns, err := b.DBCli.QueryTioServerById(sid)
	if err != nil {
		return err
	}

	ns.Status = stauts

	return b.DBCli.UpdateTioServer(ns)
}

func QueryUserAllSrv(b *data.B, uid, limit int) ([]model.Server, error) {
	return b.DBCli.QueryTioServerByUser(uid, limit)
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
