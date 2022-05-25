package routers

import (
	"Go-with-Hitesh/Mongo_API/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// Get all movies route
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	// Get a movie route
	router.HandleFunc("/api/movie/{id}", controllers.GetAMovie).Methods("GET")
	// Post a movie route
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	// Update a movie as watched route
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	// Delete a movie
	router.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	// Delete all movies
	router.HandleFunc("/api/deletemovies", controllers.DeleteAllMovies).Methods("DELETE")
	return router
}
