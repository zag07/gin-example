package server

import "context"

// ProviderSet is server providers.
// var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)

// Server is transport server.
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}