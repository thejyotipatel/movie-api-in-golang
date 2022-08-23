package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thejyotipatel/movies/pkg/controllers"
	"github.com/thejyotipatel/movies/pkg/models"
	"github.com/thejyotipatel/movies/pkg/routes"
)

func main() {
	fmt.Println("Serverrunning at port 8080...")

	r := mux.NewRouter()
	routes.RegisterRoute(r)

	controllers.NewMovie = append(controllers.NewMovie, models.Movie{Id: "1", MovieName: "movie1", ReleaseDate: "12 dec wed", Director: &models.Director{Name: "director1", Age: 23}})

	controllers.NewMovie = append(controllers.NewMovie, models.Movie{Id: "2", MovieName: "movie2", ReleaseDate: "24 jan tue", Director: &models.Director{Name: "director2", Age: 45}})

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
