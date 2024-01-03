package service

import (
	v1 "bphn/artikel-hukum/api/v1"
	"bphn/artikel-hukum/internal/errors"
	"bphn/artikel-hukum/internal/ito"
	"bphn/artikel-hukum/internal/model"
	mock_repository "bphn/artikel-hukum/internal/repository/mocks"
	"bphn/artikel-hukum/pkg/config"
	"bphn/artikel-hukum/pkg/log"
	"context"
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var (
	service Service
	logger  *log.Logger
)

func TestMain(m *testing.M) {
	err := os.Setenv("APP_CONF", "../../config/local.yml")
	if err != nil {
		fmt.Println("Setenv error", err)
	}
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)

	service = Service{logger: logger}

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserService_List(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	query := ito.ListQuery{
		Page:   1,
		Size:   10,
		Sort:   "+id",
		Filter: "email",
	}

	repo := mock_repository.NewMockUserRepository(ctrl)
	repo.EXPECT().FindAll(gomock.Any(), query).Return(&ito.ListQueryResult[model.User]{
		TotalPage: 1,
		Page:      1,
		Items: []model.User{
			{
				Id:              1,
				FullName:        "John Doen",
				Password:        "12345678",
				Email:           "mail@johndoe.com",
				EmailVerifiedAt: time.Time{},
				Role:            "author",
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				DeletedAt:       gorm.DeletedAt{},
			},
			{
				Id:              2,
				FullName:        "John Doen",
				Password:        "12345678",
				Email:           "mail2@johndoe.com",
				EmailVerifiedAt: time.Time{},
				Role:            "author",
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				DeletedAt:       gorm.DeletedAt{},
			},
		},
	}, nil).AnyTimes()

	userService := NewUserService(&service, repo)

	users, err := userService.List(context.Background(), query)

	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(users.Items))
	}
}

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_repository.NewMockUserRepository(ctrl)

	t.Run("Success create user", func(t *testing.T) {
		createUser := v1.CreateUserRequest{
			FullName: "John Doe",
			Email:    "mail@johndoe.com",
			Password: "12345678",
			Role:     "author",
		}

		gomock.InOrder(
			repo.EXPECT().FindByEmail(gomock.Any(), createUser.Email).Return(nil, nil).Times(1),
			repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1),
		)

		userService := NewUserService(&service, repo)

		err := userService.Create(context.Background(), &createUser)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, nil, err)

	})

	t.Run("Failed create user due to email already exists", func(t *testing.T) {
		createUser := v1.CreateUserRequest{
			FullName: "John Doe",
			Email:    "mail@johndoe.com",
			Password: "12345678",
			Role:     "author",
		}

		r := &model.User{
			FullName: "John Doe",
			Email:    "mail@johndoe.com",
			Password: "12345678",
			Role:     "author"}

		repo.EXPECT().FindByEmail(gomock.Any(), createUser.Email).Times(1).Return(r, nil)

		userService := NewUserService(&service, repo)

		err := userService.Create(context.Background(), &createUser)

		assert.Error(t, err)

	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		user := v1.UpdateUserRequest{
			Id:       12,
			FullName: "john doe",
			Email:    "mail@johndoe.com",
			Role:     "author",
		}

		repo := mock_repository.NewMockUserRepository(ctrl)

		repo.EXPECT().FindById(gomock.Any(), user.Id).Return(&model.User{}, nil).Times(1)
		repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		userService := NewUserService(&service, repo)

		assert.NoError(t, userService.Update(context.Background(), &user))
	})

	t.Run("Failed, user not found", func(t *testing.T) {

		ctrl := gomock.NewController(t)

		repo := mock_repository.NewMockUserRepository(ctrl)
		userService := NewUserService(&service, repo)

		user := v1.UpdateUserRequest{
			Id:       12,
			FullName: "john doe",
			Email:    "mail@johndoe.com",
			Role:     "author",
		}

		repo.EXPECT().FindById(gomock.Any(), user.Id).Return(nil, nil).Times(1)
		repo.EXPECT().Update(gomock.Any(), &model.User{}).Return(errors.ErrUserDoesNotExists).Times(0)

		err := userService.Update(context.Background(), &user)

		if err != nil {
			assert.EqualError(t, errors.ErrUserDoesNotExists, err.Error())
			return
		}

		t.Fail()

	})
}

func TestUserService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_repository.NewMockUserRepository(ctrl)
	repo.EXPECT().Delete(gomock.Any(), uint(12)).Return(nil).Times(1)

	userService := NewUserService(&service, repo)

	assert.NoError(t, userService.Delete(context.Background(), uint(12)))
}

func TestUserService_FindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_repository.NewMockUserRepository(ctrl)
	repo.EXPECT().FindByEmail(gomock.Any(), "mail@johndoe.com").Return(&model.User{
		Id:       uint(12),
		FullName: "John Doe",
	}, nil).Times(1)

	userService := NewUserService(&service, repo)

	user, err := userService.FindByEmail(context.Background(), "mail@johndoe.com")
	if assert.NoError(t, err) {
		assert.Equal(t, uint(12), user.Id)
		assert.Equal(t, "John Doe", user.FullName)
	}
}

func TestUserService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_repository.NewMockUserRepository(ctrl)
	repo.EXPECT().FindById(gomock.Any(), uint(12)).Return(&model.User{
		Id:       uint(12),
		FullName: "John Doe",
	}, nil).Times(1)

	userService := NewUserService(&service, repo)

	user, err := userService.FindById(context.Background(), 12)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(12), user.Id)
		assert.Equal(t, "John Doe", user.FullName)
	}
}

func TestUserService_ChangePasswordByNonAdmin(t *testing.T) {
	t.Run("Success change password, old password is correct", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		userIndRecord := model.User{
			Id:              12,
			FullName:        "John Doe",
			Password:        hashPassword("12345678"),
			Email:           "",
			EmailVerifiedAt: time.Time{},
			Avatar:          "",
			Role:            "",
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
			DeletedAt:       gorm.DeletedAt{},
		}

		mockRepo := mock_repository.NewMockUserRepository(ctrl)
		mockRepo.EXPECT().FindById(gomock.Any(), uint(12)).Return(&userIndRecord, nil).Times(1)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		userService := NewUserService(&service, mockRepo)

		err := userService.ChangePasswordByNonAdmin(context.Background(), ito.ChangePasswordQuery{
			UserId:   12,
			Password: "12345678",
		})

		assert.NoError(t, err)

	})

	t.Run("Failed to change password due to old password is incorrect", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		userIndRecord := model.User{
			Id:              12,
			FullName:        "John Doe",
			Password:        hashPassword("12345678"),
			Email:           "",
			EmailVerifiedAt: time.Time{},
			Avatar:          "",
			Role:            "",
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
			DeletedAt:       gorm.DeletedAt{},
		}

		mockRepo := mock_repository.NewMockUserRepository(ctrl)
		mockRepo.EXPECT().FindById(gomock.Any(), uint(12)).Return(&userIndRecord, nil).Times(1)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.ErrOldPasswordIncorrect).Times(0)

		userService := NewUserService(&service, mockRepo)

		err := userService.ChangePasswordByNonAdmin(context.Background(), ito.ChangePasswordQuery{
			UserId:   12,
			Password: "12345679",
		})

		assert.Error(t, err)

	})
}

func TestUserService_ForgotPassword(t *testing.T) {

	t.Run("Success change password", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		userRepoMock := mock_repository.NewMockUserRepository(ctrl)
		userRepoMock.EXPECT().FindByEmail(gomock.Any(), "mail@johndoe.com").Return(&model.User{}, nil).Times(1)

		svc := NewUserService(&service, userRepoMock)

		err := svc.ForgotPassword(context.Background(), v1.ForgotPasswordRequest{Email: "mail@johndoe.com"})

		assert.NoError(t, err)

	})

	t.Run("Change password while email doesn't registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		userRepoMock := mock_repository.NewMockUserRepository(ctrl)
		userRepoMock.EXPECT().FindByEmail(gomock.Any(), "mail@johndoe.com").Return(nil, nil).Times(1)

		svc := NewUserService(&service, userRepoMock)

		err := svc.ForgotPassword(context.Background(), v1.ForgotPasswordRequest{Email: "mail@johndoe.com"})

		assert.Error(t, err)
		assert.Error(t, errors.ErrUserDoesNotExists, err)

	})
}

func hashPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pass)
}
