package service

import (
	"bphn/artikel-hukum/api/v1"
	"bphn/artikel-hukum/internal/repository"
	"context"
)

type UserService interface {
	List(ctx context.Context) ([]v1.UserDataResponse, error)
	Create(ctx context.Context, request *v1.CreateUserRequest) error
	Update(ctx context.Context, request *v1.UpdateUserRequest) error
	Delete(ctx context.Context, request uint) error
}

type userService struct {
	*Service
	repository repository.UserRepository
}

func NewUserService(service *Service, userRepository repository.UserRepository) UserService {
	return &userService{
		service,
		userRepository,
	}
}

func (u userService) List(ctx context.Context) ([]v1.UserDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) Create(ctx context.Context, request *v1.CreateUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Update(ctx context.Context, request *v1.UpdateUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Delete(ctx context.Context, request uint) error {
	//TODO implement me
	panic("implement me")
}
