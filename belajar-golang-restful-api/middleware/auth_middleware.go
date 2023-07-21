package middleware

import (
	"belajar_golang_restful_api/helpers"
	"belajar_golang_restful_api/models/web"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-KEY") == "RAHASIA" {
		// OK
		middleware.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		encoder := json.NewEncoder(w)
		errWeb := encoder.Encode(webResponse)
		helpers.PanicIfError(errWeb)
	}
}
