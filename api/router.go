package api

import (
	"net/http"

	"github.com/gowesmart/api-gowesmart/app"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin := app.NewRouter()

	gin.ServeHTTP(w, r)
}
