package db

import (
	"errors"

	"tio/control/data"
	"tio/database/model"
)

// QueryUser
// 当口令匹配时,返回用户信息. 否则通过error返回错误信息
func QueryUser(b *data.B, name, passwd string) (u model.User, err error) {
	u, err = b.DBCli.QueryTioUser(name)
	if err != nil {
		return
	}

	if u.Passwd == decodePasswd(passwd) {
		return u, nil
	}
	return u, errors.New("Passwd Wrong")
}

func decodePasswd(passwd string) string {
	return passwd
}

func encodePasswd(passwd string) string {
	return passwd
}

func RegisterUser(b *data.B, name, passwd string) error {
	return b.DBCli.SaveTioUser(&model.User{
		Name:   name,
		Passwd: encodePasswd(passwd),
	})
}

func ChangePasswd(b *data.B, name, oldPasswd, newPasswd string) error {
	u, err := QueryUser(b, name, oldPasswd)
	if err != nil {
		return err
	}

	u.Passwd = encodePasswd(newPasswd)
	return b.DBCli.UpdateTioUser(&u)
}
