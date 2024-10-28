package services

import (
	"database/sql"
	"encoding/json"
	"go-movie-crud/pkg/models"
	"log"
	"net/http"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Open the database connection
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the SQL insert statement
	query := `
		INSERT INTO movies (tmdb_id, title, overview, release_date, language, popularity, vote_average, vote_count, poster_path, backdrop_path)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	// Execute the insert statement
	_, err = db.Exec(query, movie.TMDBID, movie.Title, movie.Overview, movie.ReleaseDate, movie.Language, movie.Popularity, movie.VoteAverage, movie.VoteCount, movie.PosterPath, movie.BackdropPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created movie
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(movie)
}
