package models

type Movie struct {
	ID           int    `json:"id"`
	TMDBID       int    `json:"tmdb_id"`
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	ReleaseDate  string `json:"release_date"`
	Language     string `json:"language"`
	Popularity   string `json:"popularity"`
	VoteAverage  string `json:"vote_average"`
	VoteCount    string `json:"vote_count"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Movies []Movie

type TMDBResponse struct {
	Results []TMDBMovie `json:"results"`
}
type TMDBMovie struct {
	ID           int     `json:"id"`
	TMDBID       int     `json:"tmdb_id"`
	Title        string  `json:"title"`
	Overview     string  `json:"overview"`
	ReleaseDate  string  `json:"release_date"`
	Language     string  `json:"original_language"`
	Popularity   float64 `json:"popularity"`
	VoteAverage  float64 `json:"vote_average"`
	VoteCount    int     `json:"vote_count"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
}

type Config struct {
	Apikey  string
	BaseURL string
}

var NewConfig Config
