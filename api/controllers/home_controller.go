package controllers

import (
	"net/http"

	"github.com/AlexSwiss/reev-v1/api/responses"
)

// Home function welcomes developer to the api
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To The API")
}
