package server

import "time"

type Server interface {
	Start() error
	ShutDown(killTime time.Duration) error
}
