package controllers

import (
	"net/http"
	"strconv"

	"github.com/DLzer/icf/api/models"
	"github.com/DLzer/icf/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	exercise := models.Exercise{}
	exerciseRetrieved, err := exercise.GetExercise(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, exerciseRetrieved)
}
