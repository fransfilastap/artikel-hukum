package server

import (
	"bphn/artikel-hukum/pkg/log"
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type HttpServer struct {
	echo   *echo.Echo
	host   string
	port   int
	logger *log.Logger
}

type Option func(server *HttpServer)

func NewServer(engine *echo.Echo, logger *log.Logger, opts ...Option) *HttpServer {
	s := &HttpServer{echo: engine, logger: logger}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithServerHost(host string) Option {
	return func(s *HttpServer) {
		s.host = host
	}
}

func WithServerPort(port int) Option {
	return func(s *HttpServer) {
		s.port = port
	}
}

func (h *HttpServer) Engine() *echo.Echo {
	return h.echo
}

func (h *HttpServer) Start() error {
	address := fmt.Sprintf("%s:%d", h.host, h.port)
	if err := h.echo.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
		h.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}

func (h *HttpServer) ShutDown(ctx context.Context) error {
	h.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.echo.Shutdown(ctx); err != nil {
		h.logger.Sugar().Fatal("HttpServer forced to shutdown: ", err)
	}

	h.logger.Sugar().Info("HttpServer exiting")

	return nil
}
