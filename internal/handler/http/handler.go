package http

import (
	"bphn/artikel-hukum/pkg/log"
	"github.com/spf13/viper"
)

type Handler struct {
	viper  *viper.Viper
	logger *log.Logger
}

func NewHandler(config *viper.Viper, logger *log.Logger) *Handler {
	return &Handler{config, logger}
}

func (h *Handler) Logger() *log.Logger {
	return h.logger
}
