package controllers

import "github.com/DLzer/icf/api/middleware"

// initializeRoutes declares all routes used in the application
func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/user/{id:[0-9]+}", middleware.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/workout/{id:[0-9]+}", middleware.SetMiddlewareJSON(s.GetWorkout)).Methods("GET")
	s.Router.HandleFunc("/exercise/{id:[0-9]+}", middleware.SetMiddlewareJSON(s.GetExercise)).Methods("GET")
	s.Router.HandleFunc("/tracker/{id:[0-9]+}", middleware.SetMiddlewareJSON(s.GetTracker)).Methods("GET")

}
