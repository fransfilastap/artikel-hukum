//go:build wireinject

package wire

import (
	"bphn/artikel-hukum/internal/repository"
	"bphn/artikel-hukum/internal/server"
	"bphn/artikel-hukum/pkg/app"
	"bphn/artikel-hukum/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRepository)

func newApp(migration *server.Migration) *app.App {
	return app.NewApp(app.WithName("database-migration"), app.WithServer(migration))
}

func InitializeMigration(viper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(server.NewMigration, repositorySet, newApp))
}
