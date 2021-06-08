package controllers

import (
	"net/http"
	"strconv"

	"github.com/DLzer/icf/api/models"
	"github.com/DLzer/icf/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetWorkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	workout := models.Workout{}
	workoutRetrieved, err := workout.GetWorkout(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, workoutRetrieved)
}
