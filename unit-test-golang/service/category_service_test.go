package service

import (
	"my-golang/unit-test-golang/entity"
	"my-golang/unit-test-golang/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_NotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindId", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryService_Found(t *testing.T) {

	// mock return
	category := entity.Category{
		Id:   "2",
		Name: "Golang",
	}

	// program mock
	categoryRepository.Mock.On("FindId", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
