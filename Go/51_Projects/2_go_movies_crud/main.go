package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	First_name string `json:"firstName"`
	Last_name  string `json:"lastName"`
}

var movies []Movie

func generateUniqueID() string {
	id := uuid.New()
	return id.String()
}

func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	params := mux.Vars(r)
	for _, mov := range movies {

		if mov.ID == params["id"] {
			json.NewEncoder(w).Encode(mov)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)

	var newMovie Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newMovie.ID = generateUniqueID()

	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)

	params := mux.Vars(r)
	movieID := params["id"]

	var updatedMovie Movie
	err := json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the movie ID is "1" or "2" and skip the update
	if movieID == "1" || movieID == "2" {
		http.Error(w, "Cannot update movie with ID 1 or 2", http.StatusForbidden)
		return
	}

	for index, mov := range movies {
		if mov.ID == movieID {
			movies[index] = updatedMovie
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	params := mux.Vars(r)
	movieID := params["id"]

	// Check if the movie ID is "1" or "2" and skip the deletion
	if movieID == "1" || movieID == "2" {
		http.Error(w, "Cannot delete movie with ID 1 or 2", http.StatusForbidden)
		return
	}

	for index, mov := range movies {
		if mov.ID == movieID {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	getAllMovies(w, r)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "TEST_DATA1",
		Title: "Whispers in the Shadows",
		Director: &Director{
			First_name: "Olivia",
			Last_name:  "Bennett"}})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "TEST_DATA2",
		Title: "Echoes of Eternity",
		Director: &Director{
			First_name: "Samantha",
			Last_name:  "Harper"}})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/movie/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies/movie", createMovie).Methods("POST")
	r.HandleFunc("/movies/movie/{id}", updateMovie).Methods("PUT", "PATCH")
	r.HandleFunc("/movies/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")
	log.Fatal((http.ListenAndServe(":8000", r)))
}
