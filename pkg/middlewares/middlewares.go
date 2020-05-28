package middlewares

import (
	"errors"
	"net/http"

	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/auth"
	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/responses"
	"github.com/jinzhu/gorm"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc, DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidateToken(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
