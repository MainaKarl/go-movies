package services

import (
	"database/sql"
	"go-movie-crud/pkg/models"
	"log"
)

func SaveMovieToDB(db *sql.DB, movie *models.TMDBMovie) error {
	query := `
		INSERT INTO movies (tmdb_id, title, overview, release_date, language, popularity, vote_average, vote_count, poster_path, backdrop_path)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (tmdb_id) DO NOTHING;
	`

	log.Printf("Saving movie: %+v\n", movie) // Log the movie details

	_, err := db.Exec(query, movie.ID, movie.Title, movie.Overview, movie.ReleaseDate, movie.Language, movie.Popularity, movie.VoteAverage, movie.VoteCount, movie.PosterPath, movie.BackdropPath)
	if err != nil {
		log.Printf("Error saving movie to DB: %v", err)
		return err
	}
	log.Println("Movie saved successfully!")
	return nil
}
