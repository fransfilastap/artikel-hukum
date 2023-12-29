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

type MustHaveUserTrait struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type AdminCreateUserRequest struct {
	*MustHaveUserTrait
	Password string `json:"password" validate:"required,min:"`
	Role     string `json:"role" validate:"required,oneof=admin author editor"`
}

type AuthorRegistrationRequest struct {
	*MustHaveUserTrait
	Occupation string `json:"occupation" validate:"required"`
	Company    string `json:"company" validate:"required"`
}
