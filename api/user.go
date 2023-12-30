package api

type UserDataResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
	Role     string `json:"role"`
}

// TODO add pagination traits
type UserListResponse struct {
	data *[]UserDataResponse
}

type CreateUserRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required"`
}

type UpdateUserRequest struct {
	FullName             string `json:"full_name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"omitempty,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"omitempty,min=8,eqfield=Password"`
	Role                 string `json:"role" validate:"required"`
}

type AuthorRegistrationRequest struct {
	FullName   string `json:"full_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Occupation string `json:"occupation" validate:"required"`
	Company    string `json:"company" validate:"required"`
}
