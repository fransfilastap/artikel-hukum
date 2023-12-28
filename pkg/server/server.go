package server

import (
	"context"
)

type Server interface {
	Start() error
	ShutDown(ctx context.Context) error
}
