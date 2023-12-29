package repository

import (
	"bphn/artikel-hukum/internal/domain"
	"context"
)

type userRepository struct {
	*Repository
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
}
