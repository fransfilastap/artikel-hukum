package v1

type AuthorProfileDataResponse struct {
	Id         uint   `json:"id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar,omitempty"`
	Occupation string `json:"occupation"`
	Company    string `json:"company"`
}

type AuthorRegistrationRequest struct {
	FullName   string `json:"full_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Occupation string `json:"occupation" validate:"required"`
	Company    string `json:"company" validate:"required"`
	Password   string `json:"password" validate:"required,min=8"`
}

type UpdateAuthorProfileRequest struct {
	Id         uint   `json:"id,omitempty"`
	FullName   string `json:"full_name" validate:"required"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email" validate:"required,email"`
	Occupation string `json:"occupation" validate:"required"`
	Company    string `json:"company" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}
