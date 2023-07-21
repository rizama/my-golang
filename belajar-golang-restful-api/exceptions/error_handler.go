package exceptions

import (
	"belajar_golang_restful_api/helpers"
	"belajar_golang_restful_api/models/web"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	// eksekusi salah satu saja, buat sebuah kondisi

	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		encoder := json.NewEncoder(w)
		errWeb := encoder.Encode(webResponse)
		helpers.PanicIfError(errWeb)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		encoder := json.NewEncoder(w)
		errWeb := encoder.Encode(webResponse)
		helpers.PanicIfError(errWeb)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	encoder := json.NewEncoder(w)
	errWeb := encoder.Encode(webResponse)
	helpers.PanicIfError(errWeb)
}
