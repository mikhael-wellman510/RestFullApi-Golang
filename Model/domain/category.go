package domain

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required min:0 max:20"`
}
