package services

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go-movie-crud/pkg/models"
)


func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range models.Movies {
		if item.ID == params["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			models.Movies = append(models.Movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(models.Movies)
}