package services

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go-movie-crud/pkg/models"
)

func GetMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var movie models.Movie
	json.NewEncoder(w).Encode(movie)
}

func GetMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range models.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Movie{})
}