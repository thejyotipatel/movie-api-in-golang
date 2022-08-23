package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/thejyotipatel/movies/pkg/routes"
)

func main() {
	fmt.Println("Server running at port 8080...")

	route := routes.RegisterRoute(mux.NewRouter())

	route.Handler('/')
}