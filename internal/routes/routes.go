package routes

import (
	"github.com/gorilla/mux"
	"github.com/ysdinesh31/temperature-error-logs/internal/controllers"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/temp", controllers.HandleTemperatureReading).Methods("POST")
	r.HandleFunc("/errors", controllers.GetErrors).Methods("GET")
	r.HandleFunc("/errors", controllers.DeleteErrors).Methods("DELETE")
}
