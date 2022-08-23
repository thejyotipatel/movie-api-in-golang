package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thejyotipatel/movies/pkg/models"
)

var NewMovie models.Movie

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contant-Type", "appication/json")

	json.NewEncoder(w).Encode(NewMovie)
}