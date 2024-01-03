package ito

type ListQuery struct {
	Page   int    `query:"page"`
	Size   int    `query:"size"`
	Sort   string `query:"sort"`
	Filter string `query:"filter"`
}

type ListQueryResult[T any] struct {
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
	Items     []T `json:"items"`
}

type ChangePasswordQuery struct {
	UserId   uint
	Password string
}

type UserDataResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
	Role     string `json:"role"`
}
