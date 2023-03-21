package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{"1", "1234", "Avatar", &Director{"James", "Cameron"}})
	movies = append(movies, Movie{"2", "1235", "Avatar2", &Director{"James", "Cameron"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("POST")
	r.HandleFunc("/movies", createMovie).Mehtods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Mehtods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Mehtods("DELETE")

	fmt.Printf("Started server at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
