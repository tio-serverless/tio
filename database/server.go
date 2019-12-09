package database

import (
	"errors"
	"os"
	"strings"

	"tio/database/postgresql"
)

func GetDBClient() (TioDb, error) {
	switch strings.ToLower(os.Getenv("TIO_DB")) {
	case "postgres":
		p := &postgresql.TDB_Postgres{}
		if err := p.Init(); err != nil {
			return nil, err
		}

		return p, nil
	}

	return nil, errors.New("TIO_DB is emtpy!")
}
