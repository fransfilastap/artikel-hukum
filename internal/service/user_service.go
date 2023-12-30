package service

import (
	"bphn/artikel-hukum/api"
	"bphn/artikel-hukum/internal/repository"
	"context"
)

type UserService interface {
	List(ctx context.Context) ([]api.UserDataResponse, error)
	Create(ctx context.Context, request *api.CreateUserRequest) error
	Update(ctx context.Context, request *api.UpdateUserRequest) error
	Delete(ctx context.Context, request uint) error
}

type userService struct {
	*Service
	repository repository.UserRepository
}

/*func NewUserService(service *Service, userRepository repository.UserRepository) UserService {
	return &userService{
		Service:    service,
		repository: userRepository,
	}
}

func (u *userService) List(ctx context.Context) ([]api.UserDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Create(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}*/
