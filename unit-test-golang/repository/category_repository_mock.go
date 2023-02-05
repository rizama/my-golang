package repository

import (
	"my-golang/unit-test-golang/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindId(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments[0] == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}

}
