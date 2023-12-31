package service

import (
	v1 "bphn/artikel-hukum/api/v1"
	"context"
)

type AuthorService interface {
	Register(ctx context.Context, request v1.AuthorRegistrationRequest) error
	ForgotPassword(ctx context.Context, request v1.ForgotPasswordRequest) error
	UpdateProfile(ctx context.Context, request v1.UpdateAuthorProfileRequest) error
	Profile(ctx context.Context, Id uint) (v1.AuthorProfileDataResponse, error)
}
