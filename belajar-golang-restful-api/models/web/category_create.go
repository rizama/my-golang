package web

type CategoryCreateRequest struct {
	Name string
}

type CategoryCreateResponse struct {
	Id   int
	Name string
}
