package fakes

import (
	"bphn/artikel-hukum/api"
	"context"
)

type FakeUserService struct{}

func (f *FakeUserService) Create(ctx context.Context) error {
	return nil
}

func (f *FakeUserService) List(ctx context.Context) ([]api.UserDataResponse, error) {
	var users = []api.UserDataResponse{
		{
			Id:       1,
			FullName: "Frans Filasta Pratama",
			Email:    "mail@fransfp.dev",
			Avatar:   "https://ui-avatars.com/api/?uppercase=false&name=frans",
			Role:     "admin",
		},
		{
			Id:       2,
			FullName: "Rahma Fitri",
			Email:    "rahmafitri92@gmail.com",
			Avatar:   "https://ui-avatars.com/api/?uppercase=false&name=rahma+fitri",
			Role:     "author",
		},
		{
			Id:       3,
			FullName: "Ibrahim Finra Achernar",
			Email:    "finn@fransfp.dev",
			Avatar:   "https://ui-avatars.com/api/?uppercase=false&name=finn",
			Role:     "author",
		},
	}

	return users, nil
}
