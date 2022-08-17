package router

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/draganaF/gled.io/UserService/model"
	"github.com/draganaF/gled.io/UserService/utils"
)

var Authenticate = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.JSONResponse(w, 400, "No header")
			return
		}

		fmt.Println(tokenString)
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		_, err := utils.VerifyJwtToken(tokenString)
		if err != nil {
			utils.JSONResponse(w, 401, "You are not authorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}

var Authorize = func(next http.Handler, roles ...model.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		role := r.Header.Get("role")
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
