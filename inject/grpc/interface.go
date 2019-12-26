package main

//go:generate mockgen -destination mock/injectGrpcMock.go -package mock -source interface.go
type injectGrpc interface {
	FetchServices(string) ([]string, error)
	FetchMethods(string, string) ([]string, error)
	Store(name string, methods []string) error
}
