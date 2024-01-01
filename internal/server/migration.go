package server

import (
	"bphn/artikel-hukum/internal/model"
	"bphn/artikel-hukum/pkg/log"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type Migration struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewMigration(db *gorm.DB, logger *log.Logger) *Migration {
	return &Migration{db: db, logger: logger}
}

func (m *Migration) Start() error {
	if err := m.db.AutoMigrate(&model.User{}, &model.AuthorDetail{}, &model.PasswordResetToken{}); err != nil {
		m.logger.Error("migration error", zap.Error(err))
		panic(err)
	}
	m.logger.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}

func (m *Migration) ShutDown(ctx context.Context) error {
	m.logger.Info("Automigrate stop")
	return nil
}
