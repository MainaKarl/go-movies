package main

import (
	"database/sql"
	"fmt"
	"go-movie-crud/pkg/models"
	"go-movie-crud/pkg/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = models.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	//db connection
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println(db)

	// Example of fetching and saving a movie from TMDB
	movie, err := services.FetchMovieFromTMDB(db)
	if err != nil {
		log.Fatal("Error fetching movie:", err)
	}

	log.Printf("Fetched movie: %+v\n", movie)

	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/movies", services.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", services.GetMovie).Methods("GET")
	r.HandleFunc("/api/movies", services.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", services.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", services.DeleteMovie).Methods("DELETE")

	// Start server
	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
