package controllers

import (
	"belajar_golang_restful_api/helpers"
	"belajar_golang_restful_api/models/web"
	"belajar_golang_restful_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helpers.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	errWeb := encoder.Encode(webResponse)
	helpers.PanicIfError(errWeb)
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helpers.PanicIfError(err)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	// set & update id to int
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	errWeb := encoder.Encode(webResponse)
	helpers.PanicIfError(errWeb)
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	errWeb := encoder.Encode(webResponse)
	helpers.PanicIfError(errWeb)
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	categoryFindResponse := controller.CategoryService.FindById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryFindResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	errWeb := encoder.Encode(webResponse)
	helpers.PanicIfError(errWeb)

}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categories := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categories,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	errWeb := encoder.Encode(webResponse)
	helpers.PanicIfError(errWeb)
}
