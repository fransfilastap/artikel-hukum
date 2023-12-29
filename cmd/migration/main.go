package main

import (
	"bphn/artikel-hukum/cmd/migration/wire"
	"bphn/artikel-hukum/pkg/config"
	"bphn/artikel-hukum/pkg/log"
	"context"
	"flag"
)

func main() {
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.InitializeMigration(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
