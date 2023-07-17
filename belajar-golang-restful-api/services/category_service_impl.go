package services

import (
	"belajar_golang_restful_api/helpers"
	"belajar_golang_restful_api/models/domain"
	"belajar_golang_restful_api/models/web"
	"belajar_golang_restful_api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateResponse {
	// Validation
	errValidation := service.Validate.Struct(request)
	helpers.PanicIfError(errValidation)

	// Using transaction when save to db, wrap all in tx
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	// Apply Category Name
	category := domain.Category{
		Name: request.Name,
	}

	// Save using repository
	category = service.CategoryRepository.Save(ctx, tx, category)

	return helpers.ToCategoryCreateResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateResponse {
	// Validation
	errValidation := service.Validate.Struct(request)
	helpers.PanicIfError(errValidation)

	// Using transaction when save to db, wrap all in tx
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	// Find Category
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helpers.PanicIfError(err)

	// Apply Category Name
	category.Name = request.Name

	// Update using repository
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helpers.ToCategoryUpdateResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// Using transaction when save to db, wrap all in tx
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	// Find Category
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helpers.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryFindResponse {
	// Using transaction when save to db, wrap all in tx
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	// Find Category
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helpers.PanicIfError(err)

	return helpers.ToCategoryFindResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryFindResponse {
	// Using transaction when save to db, wrap all in tx
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	// Find Categories
	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helpers.ToCategoresFindResponse(categories)
}
