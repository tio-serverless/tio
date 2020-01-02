package database

import "tio/database/model"

type TioDb interface {
	Init(string) error
	Version() string

	SaveTioUser(user *model.User) error
	QueryTioUser(string) (model.User, error)
	UpdateTioUser(user *model.User) error
	DeleteTioUser(string) error

	SaveTioServer(server *model.Server) error
	QueryTioServer()([]model.Server, error)
	QueryTioServerByUser(int, int, string) ([]model.Server, error)
	QueryTioServerByName(string) (*model.Server, error)
	QueryTioServerById(int) (*model.Server, error)
	UpdateTioServer(server *model.Server) error
	DeleteTioServer(string) error
}
