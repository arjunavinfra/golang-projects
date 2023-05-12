package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Director struct {
	Fristname string `json:"firstname"`
	LastName  string `json:"lastname"`
}
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"lsbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, content := range movies {
		if params["id"] == content.ID {
			json.NewEncoder(w).Encode(content)
		}
	}
}

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "3244", Title: "IronMan", Director: &Director{Fristname: "john", LastName: "lucas"}})
	movies = append(movies, Movie{ID: "2", Isbn: "6543", Title: "Spiderman", Director: &Director{Fristname: "john", LastName: "lucas"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies/", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
