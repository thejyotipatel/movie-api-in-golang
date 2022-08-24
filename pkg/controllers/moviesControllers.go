package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/thejyotipatel/movies/pkg/models"
)

var NewMovie []models.Movie

// display on browser
func Serverpage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<h1> welcome to movie API</h1>
	`))
}

// get all data
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all movie")

	w.Header().Set("Contant-Type", "appication/json")

	json.NewEncoder(w).Encode(NewMovie)
}

// get one data by id
// 1. get ID from url
// 2. loop through all data
// 3. check loop id and param id is maches or not
// 4. if mach then return
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get data by id")

	w.Header().Set("Contant-Type", "application/json")

	// get id
	param := mux.Vars(r)

	for _, item := range NewMovie {
		if item.Id == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&NewMovie)
}

// add data
// 1. check if body is empty or nil
// 2. create variable of struct to append data
// 3. append data to this variable
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("created data")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		http.Error(w, "please send some date", http.StatusBadRequest)
		return
	} else {

		var m models.Movie
		_ = json.NewDecoder(r.Body).Decode(&m)

		// fmt.Printf("%T", r.Body)
		rand.Seed(time.Now().UnixNano())
		m.Id = string(rand.Intn(100))

		NewMovie = append(NewMovie, m)

		json.NewEncoder(w).Encode(m)
	}

}

// edit data
// 1. get id from url
// 2. loop through all data
// 3. check loop id and param id is maches or not
// 4. if maches update that index value append
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update data")

	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for index, m := range NewMovie {
		if m.Id == param["id"] {
			NewMovie = append(NewMovie[:index], NewMovie[index+1:]...)

			var m models.Movie
			_ = json.NewDecoder(r.Body).Decode(&m)
			m.Id = param["id"]

			NewMovie = append(NewMovie, m)
			json.NewEncoder(w).Encode(m)
			return
		}
	}
}

// delete data
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete data")

	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for index, m := range NewMovie {
		if m.Id == param["id"] {
			NewMovie = append(NewMovie[:index], NewMovie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(&NewMovie)

}
