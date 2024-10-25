package services

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"go-movie-crud/pkg/models"
)


func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	models.Movies = append(models.Movies, movie)
	json.NewEncoder(w).Encode(movie)
}