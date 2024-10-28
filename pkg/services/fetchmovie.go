package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-movie-crud/pkg/models"
	"net/http"
)

func FetchMovieFromTMDB(db *sql.DB) (*models.TMDBMovie, error) {
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

	// Check if we received any results
	// if len(tmdbResponse.Results) == 0 {
	// 	return nil, fmt.Errorf("no results found for movie ID %s", movieID)
	// }

	movie := tmdbResponse.Results[0] // Get the first result

	// Save fetched movie to the database
	if err := SaveMovieToDB(db, &movie); err != nil {
		return nil, err
	}

	return &movie, nil
}
