package router

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/draganaF/gled.io/UserService/model"
	"github.com/draganaF/gled.io/UserService/utils"
)

var Authenticate = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		_, err := ReturnClaims(tokenString)

		if err != nil && err.Error() == "No header found" {
			utils.JSONResponse(w, 400, err)
			return
		}

		if err != nil && err.Error() == "You are not authorized" {
			utils.JSONResponse(w, 401, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var Authorize = func(next http.Handler, roles ...model.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		claims, err := ReturnClaims(tokenString)

		if err != nil && err.Error() == "No header found" {
			utils.JSONResponse(w, 400, err)
			return
		}

		if err != nil && err.Error() == "You are not authorized" {
			utils.JSONResponse(w, 401, err)
			return
		}

		role := claims.(jwt.MapClaims)["role"].(string)

		contains := false
		for _, ar := range roles {
			if role == strconv.Itoa((int(ar))) {
				contains = true
				break
			}
		}

		if contains {
			next.ServeHTTP(w, r)
		} else {
			utils.JSONResponse(w, 403, "You don't have permission")
		}
	})
}

func ReturnClaims(tokenString string) (jwt.Claims, error) {
	if len(tokenString) == 0 {
		return nil, errors.New("No header")
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := utils.VerifyJwtToken(tokenString)

	if err != nil {
		return nil, errors.New("You are not authorized")
	}

	return claims, nil
}
