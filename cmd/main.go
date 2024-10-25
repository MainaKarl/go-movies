package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-movie-crud/pkg/services"
	"go-movie-crud/pkg/models"
)


func main(){
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	models.Movies = append(models.Movies, models.Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &models.Director{Firstname: "John", Lastname: "Doe"}})
	models.Movies = append(models.Movies, models.Movie{ID: "2", Isbn: "847564", Title: "Movie Two", Director: &models.Director{Firstname: "Steve", Lastname: "Smith"}})

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

