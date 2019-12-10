package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"tio/database/model"
)

type TDB_Postgres struct {
	db *sql.DB
}

func (p *TDB_Postgres) Init(addr string) error {
	//addr := os.Getenv("TIO_DB_POSTGRES_CONN")
	logrus.Debugf("Postgres Connstr: %s", addr)

	db, err := sql.Open("postgres", addr)
	if err != nil {
		logrus.Fatalf("Connect Postgres Error: %s", err.Error())
	}

	p.db = db

	if err := p.db.Ping(); err != nil {
		logrus.Fatalf("Ping Postgres Error: %s", err.Error())
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(5 * time.Minute)

	logrus.Info("Connect Postgres Success.")

	return nil
}

func (p *TDB_Postgres) Version() string {
	return "PosgresQL"
}

func (p *TDB_Postgres) SaveTioUser(user *model.User) error {
	sql := "INSERT INTO user(name, passwd) VALUES ($1, $2)"
	logrus.Debugf("Save New User: [%s]", sql)

	_, err := p.db.Exec(sql, user.Name, user.Passwd)
	return err
}

func (p *TDB_Postgres) QueryTioUser(name string) (model.User, error) {

	u := model.User{}

	sql := "SELECT * FROM user WHERE name=$1"
	logrus.Debugf("Query User: [%s]", sql)
	rows, err := p.db.Query(sql, name)
	if err != nil {
		return u, err
	}

	if rows.Next() {
		err = rows.Scan(&u.Id, &u.Name, &u.Passwd)
	} else {
		return u, errors.New("No Match User Record")
	}

	return u, nil
}
func (p *TDB_Postgres) UpdateTioUser(user *model.User) error {
	sql := "UPDATE user SET passwd=$2 WHERE name=$1"
	logrus.Debugf("Update User: [%s]", sql)

	_, err := p.db.Exec(sql, user.Name, user.Passwd)

	return err
}

func (p *TDB_Postgres) DeleteTioUser(name string) error {
	sql := "DELETE user WHERE name=$1"
	logrus.Debugf("Delete User: [%s]", sql)

	_, err := p.db.Exec(sql, name)

	return err
}

func (p *TDB_Postgres) SaveTioServer(s *model.Server) error {
	sql := "INSERT INTO server (name, version, uid, stype, domain, path, tversion, timestamp, status, image, raw) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	logrus.Debugf("Save New Server:[%s]", sql)

	_, err := p.db.Exec(sql, s.Name, s.Version, s.Uid, s.Stype, s.Domain, s.Path, s.TVersion, s.Timestamp, s.Status, s.Image, s.Raw)
	return err
}

func (p *TDB_Postgres) QueryTioServerByUser(uid, limit int) ([]model.Server, error) {
	var ss []model.Server

	var sql string
	if limit > 0 {
		sql = fmt.Sprintf("SELECT * FROM server WHERE uid=$1 ORDER BY timestamp desc LIMIT %d", limit)
	} else {
		sql = fmt.Sprintf("SELECT * FROM server WHERE uid=$1 ORDER BY timestamp desc")
	}

	logrus.Debugf("Query Server:[%s]", sql)
	rows, err := p.db.Query(sql, uid)
	if err != nil {
		return ss, err
	}

	for rows.Next() {
		s := model.Server{}
		err = rows.Scan(&s.Id, &s.Name, &s.Version, &s.Uid, &s.Stype, &s.Domain, &s.Path, &s.TVersion, &s.Timestamp, &s.Status, &s.Image, &s.Raw)
		if err != nil {
			logrus.Errorf("Scan Server Error. %s", err)
			continue
		}
		ss = append(ss, s)
	}

	return ss, nil
}

func (p *TDB_Postgres) QueryTioServerById(sid int) (*model.Server, error) {
	sql := fmt.Sprintf("SELECT * FROM server WHERE id=$1")
	logrus.Debugf("Query Server:[%s] id: [%d]", sql, sid)
	s := model.Server{}

	rows, err := p.db.Query(sql, sid)
	if err != nil {
		return &s, err
	}

	for rows.Next() {
		err = rows.Scan(&s.Id, &s.Name, &s.Version, &s.Uid, &s.Stype, &s.Domain, &s.Path, &s.TVersion, &s.Timestamp, &s.Status, &s.Image, &s.Raw)
		if err != nil {
			return nil, err
		}
	}

	return &s, nil
}

func (p *TDB_Postgres) QueryTioServerByName(name string) (*model.Server, error) {
	sql := fmt.Sprintf("SELECT * FROM server WHERE name=$1")
	logrus.Debugf("Query Server:[%s]", sql)
	s := model.Server{}

	rows, err := p.db.Query(sql, name)
	if err != nil {
		return &s, err
	}

	if rows.Next() {
		err = rows.Scan(&s.Id, &s.Name, &s.Version, &s.Uid, &s.Stype, &s.Domain, &s.Path, &s.TVersion, &s.Timestamp, &s.Status, &s.Image, &s.Raw)
	} else {
		return &s, errors.New("No Match Server")
	}

	return &s, nil
}

func (p *TDB_Postgres) UpdateTioServer(s *model.Server) error {
	sql := "UPDATE server SET version=$2, stype=$3, domain=$4, path=$5, tversion=$6, timestamp=$7, status=$8,image=$9, raw=$10 WHERE name=$1"
	logrus.Debugf("Update Server: [%s]", sql)

	_, err := p.db.Exec(sql, s.Name, s.Version, s.Stype, s.Domain, s.Path, s.TVersion, s.Timestamp, s.Status, s.Image, s.Raw)

	return err
}

func (p *TDB_Postgres) DeleteTioServer(name string) error {
	sql := "DELETE server WHERE name=$1"
	logrus.Debugf("Delete server: [%s]", sql)

	_, err := p.db.Exec(sql, name)

	return err

}
