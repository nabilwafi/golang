package web

type CategoryUpdateRequest struct {
	Id int `validate:"required" json:"id"`
	Name string `validate:"required,max=100" json:"name"`
}