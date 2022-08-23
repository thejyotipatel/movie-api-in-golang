package routes

import (
	"github.com/gorilla/mux"
	"github.com/thejyotipatel/movies/pkg/controllers"
)

var RegisterRoute = func(route *mux.Router){
	
	route.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	// route.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	// route.HandleFunc("/api/movies/{id}", controllers.GetMovieById).Methods("GET")
	// route.HandleFunc("/api/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	// route.HandleFunc("/api/movies/{id}", controllers.DeleteMovie).Methods("DELETE")

}
