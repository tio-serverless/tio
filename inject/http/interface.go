package main

//go:generate mockgen -destination mock/injectHttpMock.go -package mock -source interface.go
type injectHttp interface {
	Store(name string, urls []string) error
}
