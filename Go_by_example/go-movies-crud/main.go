package main

import ("fmt"
		"net/http"
		"github.com/gorilla/mux"
		"math/rand"
		"strconv"
		"encoding/json"
		"log"
)

type Movie struct {
		ID string `json:"id"`
		Isbn string `json:"isbn"`
		Title string `json:"title"`
		Director *Director `json:"director"`
}

type Director struct {
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
}

var movies [] Movie    //slice of type movie

func getMovies(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(movies)
}

func deleteMovie (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for item, index := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
}

func main () {
	r := mux.NewRouter()

	movies = append(movies, Movies{ID: "1", Isbn: "123456789", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "14578686", Title: "Movie Two", Director: &Director{Firstname: "Steven", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "15265856", Title: "Movie Three", Director: &Director{Firstname: "Earl", Lastname: "Mogire"}})
	movies = append(movies, Movie{ID: "4", Isbn: "56565562", Title: "Movie Four", Director: &Director{Firstname: "Spencer", Lastname: "Earl"}})
	movies = append (movies, Movie{ID: "5", Isbn: "56586258", Title: "Movie Five", Director: &Director{Firstname: "Most", Lastname: "Excellent"}})
	//need five different routes and five different functions -- see draw.io file
	r.HandleFunc("/movies", getMovies).Methods("GET") 
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}