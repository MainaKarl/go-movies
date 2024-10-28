package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-movie-crud/pkg/models"
	"net/http"
)

func FetchMovieFromTMDB(db *sql.DB) ([]models.TMDBMovie, error) {
	apiKey := models.NewConfig.Apikey
	baseUrl := models.NewConfig.BaseURL
	url := fmt.Sprintf("%s?api_key=%s", baseUrl, apiKey)

	fmt.Println("Fetching movie from URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tmdbResponse models.TMDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResponse); err != nil {
		return nil, err
	}

	for _, movie := range tmdbResponse.Results {
		if err := SaveMovieToDB(db, &movie); err != nil {
			fmt.Printf("Error saving movie %s to DB: %v\n", movie.Title, err)
		}
	}

	return tmdbResponse.Results, nil
}
