package service

import (
	v1 "bphn/artikel-hukum/api/v1"
	"bphn/artikel-hukum/internal/dto"
	"bphn/artikel-hukum/internal/model"
	mock_repository "bphn/artikel-hukum/internal/repository/mocks"
	"bphn/artikel-hukum/pkg/config"
	"bphn/artikel-hukum/pkg/log"
	"context"
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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

	service = Service{logger: nil}

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserService_List(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	query := dto.ListQuery{
		Page:   1,
		Size:   10,
		Sort:   "+id",
		Filter: "email",
	}

	repo := mock_repository.NewMockUserRepository(ctrl)
	repo.EXPECT().FindAll(gomock.Any(), query).Return(&dto.ListQueryResult[model.User]{
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

/*func TestUserService_Update(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		user := v1.UpdateUserRequest{
			Id:       12,
			FullName: "john doe",
			Email:    "mail@johndoe.com",
			Password: "12345678",
			Role:     "author",
		}

		repo := mock_repository.NewMockUserRepository(ctrl)
		repo.EXPECT().Update(gomock.Any(), &user).Return(nil).AnyTimes()

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
			Password: "12345678",
			Role:     "author",
		}

		repo.EXPECT().Update(gomock.Any(), &user).Return(v1.ErrUserDoesNotExists).AnyTimes()

		err := userService.Update(context.Background(), &user)

		if err != nil {
			assert.EqualError(t, v1.ErrUserDoesNotExists, err.Error())
			return
		}

		t.Fail()

	})
}*/
