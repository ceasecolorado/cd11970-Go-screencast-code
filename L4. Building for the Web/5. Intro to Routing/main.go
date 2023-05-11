package main

// The encoding/json package helps with the encoding and decoding of JSON
// i.e., the package allows us to convert Go values to and from JSON
import (
	"encoding/json"
	"fmt"
	"net/http"

	// Import the third-party gorilla/mux package
	"github.com/gorilla/mux"
)

// A map with string keys pointing to string values
var dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

// This handler function simply returns a JSON version of the "dictionary" map
func getDictionary(w http.ResponseWriter, r *http.Request) {
	// Sets up the Content-Type header so that the client knows to expect a JSON response
	w.Header().Set("Content-Type", "application/json")
	// write the header status code to 200 (OK)
	w.WriteHeader(http.StatusOK)

	// We cannot simply response with the "dictionary" map as-is, since the "dictionary" is a Go value
	// Instead, webuilds a new encoder using the ResponseWriter, then encodes the map into JSON
	json.NewEncoder(w).Encode(dictionary)
}

func main() {
	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	// Make a GET request to "/" calls the "getDictionary" handler, which returns the entire JSON-encoded map
	router.HandleFunc("/dictionary", getDictionary).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
