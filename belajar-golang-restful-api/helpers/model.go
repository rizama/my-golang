package helpers

import (
	"belajar_golang_restful_api/models/domain"
	"belajar_golang_restful_api/models/web"
)

func ToCategoryCreateResponse(category domain.Category) web.CategoryCreateResponse {
	return web.CategoryCreateResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryUpdateResponse(category domain.Category) web.CategoryUpdateResponse {
	return web.CategoryUpdateResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryFindResponse(category domain.Category) web.CategoryFindResponse {
	return web.CategoryFindResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoresFindResponse(categories []domain.Category) []web.CategoryFindResponse {
	var categoriesResponse []web.CategoryFindResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, web.CategoryFindResponse(category))
	}

	return categoriesResponse
}
