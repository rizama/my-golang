package services

import (
	"belajar_golang_restful_api/models/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryFindResponse
	FindAll(ctx context.Context) []web.CategoryFindResponse
}
