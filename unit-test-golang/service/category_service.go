package service

import (
	"errors"
	"my-golang/unit-test-golang/entity"
	"my-golang/unit-test-golang/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindId(id)

	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}
}
