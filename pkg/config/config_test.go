package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	t.Run("Get database host", func(t *testing.T) {
		cfg := NewConfig("../../config/local.yml")
		assert.Equal(t, cfg.Get("data.mysql.host"), "127.0.0.1")
	})
}
