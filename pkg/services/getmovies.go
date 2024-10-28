package services

import (
	"database/sql"
	"encoding/json"
	"go-movie-crud/pkg/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM movies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.TMDBID, &movie.Title, &movie.Overview, &movie.ReleaseDate, &movie.Language, &movie.Popularity, &movie.VoteAverage, &movie.VoteCount, &movie.PosterPath, &movie.BackdropPath, &movie.CreatedAt, &movie.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}

	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var movie models.Movie

	err = db.QueryRow("SELECT * FROM movies WHERE id = ?", params["id"]).Scan(&movie.ID, &movie.TMDBID, &movie.Title, &movie.Overview, &movie.ReleaseDate, &movie.Language, &movie.Popularity, &movie.VoteAverage, &movie.VoteCount, &movie.PosterPath, &movie.BackdropPath, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no movie is found, return an empty movie object
			json.NewEncoder(w).Encode(models.Movie{})
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movie)
}
