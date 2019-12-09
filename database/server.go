package database

import (
	"errors"

	"tio/database/postgresql"
)

func GetDBClient(engine, connect string) (TioDb, error) {
	switch engine {
	case "postgres":
		p := &postgresql.TDB_Postgres{}
		if err := p.Init(connect); err != nil {
			return nil, err
		}

		return p, nil
	}

	return nil, errors.New("TIO_DB is emtpy!")
}
