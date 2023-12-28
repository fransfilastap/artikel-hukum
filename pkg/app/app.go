package app

import (
	"bphn/artikel-hukum/pkg/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	name    string
	servers []server.Server
}

type Option func(app *App)

func NewApp(opts ...Option) *App {
	a := &App{}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func WithServer(servers ...server.Server) Option {
	return func(app *App) {
		app.servers = servers
	}
}

func WithName(name string) Option {
	return func(app *App) {
		app.name = name
	}
}

func (a *App) Run(ctx context.Context) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	for _, srv := range a.servers {
		go func(srv server.Server) {
			err := srv.Start()
			if err != nil {
				log.Printf("Server start err: %v", err)
			}
		}(srv)
	}

	select {
	case <-signals:
		// Received termination signal
		log.Println("Received termination signal")
	case <-ctx.Done():
		// Context canceled
		log.Println("Context canceled")
	}

	// Gracefully stop the servers
	for _, srv := range a.servers {
		err := srv.ShutDown(ctx)
		if err != nil {
			log.Printf("Server stop err: %v", err)
		}
	}

	return nil
}
