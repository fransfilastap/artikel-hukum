//go:build wireinject
// +build wireinject

package wire

import (
	"bphn/artikel-hukum/internal/server"
	"bphn/artikel-hukum/pkg/app"
	"bphn/artikel-hukum/pkg/log"
	pkgserver "bphn/artikel-hukum/pkg/server"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var serverSet = wire.NewSet(server.NewHttpServer)

func newServer(httpServer *pkgserver.HttpServer) *app.App {
	return app.NewApp(app.WithServer(httpServer), app.WithName("artikel-hukum-api"))
}

func InitializeApp(viper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(serverSet, newServer))
}
