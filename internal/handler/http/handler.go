package http

import (
	"bphn/artikel-hukum/pkg/log"
	"github.com/spf13/viper"
)

type Handler struct {
	viper  *viper.Viper
	Logger *log.Logger
}
