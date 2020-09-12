/**@Jhon Anthony Ojeda M.
    Creacion de una API Basica
    Paralelismo, concurrencia
    y sistemas distribuidos
***/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GET -->   http:localhost:8080/movies
func allMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All movies")
}

// GET -->   http:localhost:8080/movies/{id}
func findMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "find movie id")
}

// POST -->   http:localhost:8080/movies
func createMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create movie")
}

// PUT -->   http:localhost:8080/movies
func updateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update movie ")
}

// DELETE  -->   http:localhost:8080/movies/{id}
func deleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete movie id")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", allMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies/{id}", findMovieEndPoint).Methods("GET")
	r.HandleFunc("/movies", createMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", updateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovieEndPoint).Methods("DELETE")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
