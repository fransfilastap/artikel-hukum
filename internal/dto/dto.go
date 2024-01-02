package dto

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
