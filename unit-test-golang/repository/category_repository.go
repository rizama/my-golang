package repository

import "my-golang/unit-test-golang/entity"

type CategoryRepository interface {
	FindId(id string) *entity.Category
}
