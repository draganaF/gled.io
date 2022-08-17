package router

import (
	"net/http"
	"strings"

	"github.com/draganaF/gled.io/ProjectionService/utils"
)

var Authenticate = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.JSONResponse(w, 400, "No header")
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		_, err := utils.VerifyJwtToken(tokenString)
		if err != nil {
			utils.JSONResponse(w, 401, "You are not authorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
