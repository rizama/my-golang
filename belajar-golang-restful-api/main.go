package main

import (
	"belajar_golang_restful_api/app"
	"belajar_golang_restful_api/controllers"
	"belajar_golang_restful_api/helpers"
	"belajar_golang_restful_api/middleware"
	"belajar_golang_restful_api/repository"
	"belajar_golang_restful_api/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)

}
