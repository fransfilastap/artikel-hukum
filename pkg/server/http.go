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
	server *http.Server
	host   string
	port   int
	logger *log.Logger
}

type Option func(server *HttpServer)

func New(engine *echo.Echo, server *http.Server, logger *log.Logger, opts ...Option) *HttpServer {
	s := &HttpServer{echo: engine, server: server}

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

func (h *HttpServer) Start() error {
	h.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", h.host, h.port),
		Handler: h.echo,
	}

	if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		h.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}

func (h *HttpServer) ShutDown(killTime time.Duration) error {
	h.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.server.Shutdown(ctx); err != nil {
		h.logger.Sugar().Fatal("Server forced to shutdown: ", err)
	}

	h.logger.Sugar().Info("Server exiting")

	return nil
}
