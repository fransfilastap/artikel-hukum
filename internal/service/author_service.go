package service

import (
	v1 "bphn/artikel-hukum/api/v1"
	"bphn/artikel-hukum/internal/ito"
	"context"
)

type AuthorService interface {
	List(ctx context.Context, query ito.ListQuery) (ito.ListQueryResult[v1.AuthorProfileDataResponse], error)
	Register(ctx context.Context, request v1.AuthorRegistrationRequest) error
	UpdateProfile(ctx context.Context, request v1.UpdateAuthorProfileRequest) error
	Profile(ctx context.Context, Id uint) (v1.AuthorProfileDataResponse, error)
}

type authorService struct {
	*Service
}

func NewAuthorService(service *Service) AuthorService {
	return &authorService{service}
}

func (a *authorService) List(ctx context.Context, query ito.ListQuery) (ito.ListQueryResult[v1.AuthorProfileDataResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorService) Register(ctx context.Context, request v1.AuthorRegistrationRequest) error {
	//TODO implement me
	panic("implement me")
}

func (a *authorService) UpdateProfile(ctx context.Context, request v1.UpdateAuthorProfileRequest) error {
	//TODO implement me
	panic("implement me")
}

func (a *authorService) Profile(ctx context.Context, Id uint) (v1.AuthorProfileDataResponse, error) {
	//TODO implement me
	panic("implement me")
}
