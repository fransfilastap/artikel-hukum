//go:build wireinject
// +build wireinject

package wire

import (
	"bphn/artikel-hukum/internal/server"
	"bphn/artikel-hukum/pkg/app"
	"bphn/artikel-hukum/pkg/log"
	"bphn/artikel-hukum/pkg/server/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var serverSet = wire.NewSet(server.NewHttpServer)

func newServer(httpServer *http.Server) *app.App {
	return app.NewApp(app.WithServer(httpServer), app.WithName("artikel-hukum-api"))
}

func InitializeApp(viper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(serverSet, newServer))
}
