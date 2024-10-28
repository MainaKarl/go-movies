package services

import (
	"database/sql"
	"encoding/json"
	"go-movie-crud/pkg/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Open the database connection
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Decode the request body into a movie struct
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Prepare the SQL update statement
	query := `
		UPDATE movies 
		SET title = ?, overview = ?, release_date = ?, language = ?, popularity = ?, vote_average = ?, vote_count = ?, poster_path = ?, backdrop_path = ? 
		WHERE id = ?
	`

	// Execute the update statement
	result, err := db.Exec(query, movie.Title, movie.Overview, movie.ReleaseDate, movie.Language, movie.Popularity, movie.VoteAverage, movie.VoteCount, movie.PosterPath, movie.BackdropPath, params["id"])
	if err != nil {
		http.Error(w, "Failed to update movie in database", http.StatusInternalServerError)
		log.Printf("Error executing update: %v", err)
		return
	}

	// Check if the update affected any rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to verify update", http.StatusInternalServerError)
		log.Printf("Error checking rows affected: %v", err)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Movie not found", http.StatusNotFound)
		log.Printf("No movie found with id %s", params["id"])
		return
	}

	// Respond with the updated movie data
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
	log.Printf("Movie with id %s updated successfully", params["id"])
}
