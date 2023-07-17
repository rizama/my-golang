package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=100,min=1"`
}

type CategoryUpdateResponse struct {
	Id   int
	Name string
}
