package repository

import (
	"bphn/artikel-hukum/internal/ito"
	"bphn/artikel-hukum/internal/model"
	"context"
)

type authorRepository struct {
	*Repository
}

type AuthorRepository interface {
	Create(ctx context.Context, detail model.AuthorDetail) error
	Update(ctx context.Context, detail model.AuthorDetail) error
	FindById(ctx context.Context, id uint) (*model.AuthorDetail, error)
	FindAll(ctx context.Context, query ito.ListQuery) (ito.ListQueryResult[model.AuthorDetail], error)
	Delete(ctx context.Context, id uint) error
}
