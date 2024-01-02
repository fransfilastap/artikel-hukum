package service

import (
	"bphn/artikel-hukum/api/v1"
	"bphn/artikel-hukum/internal/dto"
	"bphn/artikel-hukum/internal/model"
	"bphn/artikel-hukum/internal/repository"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	List(ctx context.Context, query dto.ListQuery) (*dto.ListQueryResult[v1.UserDataResponse], error)
	FindById(ctx context.Context, id uint) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
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

func (u *userService) List(ctx context.Context, query dto.ListQuery) (*dto.ListQueryResult[v1.UserDataResponse], error) {
	users, err := u.repository.FindAll(ctx, query)
	if err != nil {
		return nil, err
	}

	userDataResponses := make([]v1.UserDataResponse, 0)

	for _, user := range users.Items {
		userDataResponses = append(userDataResponses, mapToUserDataResponse(user))
	}

	return &dto.ListQueryResult[v1.UserDataResponse]{
		TotalPage: users.TotalPage,
		Page:      users.Page,
		Items:     userDataResponses,
	}, nil
}

func (u *userService) Create(ctx context.Context, request *v1.CreateUserRequest) error {

	user, err := u.repository.FindByEmail(ctx, request.Email)

	if err != nil {
		return err
	}

	if user != nil {
		return v1.ErrEmailAlreadyExists
	}

	fmt.Println("user")

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	request.Password = string(password)
	mUser := mapRequestToModel(request)

	if err := u.repository.Create(ctx, mUser); err != nil {
		return err
	}

	return nil

}

func (u *userService) Update(ctx context.Context, request *v1.UpdateUserRequest) error {

	if user, err := u.repository.FindById(ctx, request.Id); err == nil && user == nil {
		return v1.ErrUserDoesNotExists
	} else {
		user.FullName = request.FullName
		user.Email = request.Email
		user.UpdatedAt = time.Now()
		user.Role = model.Role(request.Role)
		user.Avatar = request.Avatar

		if err := u.repository.Update(ctx, user); err != nil {
			return err
		}
	}

	return nil
}

func (u *userService) Delete(ctx context.Context, request uint) error {
	return u.repository.Delete(ctx, request)
}

func (u *userService) FindById(ctx context.Context, id uint) (*model.User, error) {
	return u.repository.FindById(ctx, id)
}

func (u *userService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.repository.FindByEmail(ctx, email)
}

func mapToUserDataResponse(user model.User) v1.UserDataResponse {
	return v1.UserDataResponse{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Role:     string(user.Role),
	}
}

func generatePasswordHash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func mapRequestToModel(request *v1.CreateUserRequest) *model.User {
	hashed, _ := generatePasswordHash(request.Password)
	return &model.User{
		FullName: request.FullName,
		Password: hashed,
		Email:    request.Email,
		Role:     model.Role(request.Role),
	}
}
