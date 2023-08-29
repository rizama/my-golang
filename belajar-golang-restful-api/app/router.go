package app

import (
	"belajar_golang_restful_api/controllers"
	"belajar_golang_restful_api/exceptions"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controllers.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// error handler
	router.PanicHandler = exceptions.ErrorHandler

	return router
}
