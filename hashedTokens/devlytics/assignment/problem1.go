package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new mux router
	router := mux.NewRouter()

	// Define endpoint-one
	router.HandleFunc("/endpoint-one", func(w http.ResponseWriter, r *http.Request) {
		// Write a response to the client
		fmt.Fprint(w, "This is endpoint one!")
	})

	// Define endpoint-two
	router.HandleFunc("/endpoint-two/{input}", func(w http.ResponseWriter, r *http.Request) {
		// Get the input parameter from the request URL
		vars := mux.Vars(r)
		input := vars["input"]

		// Write a response to the client
		fmt.Fprintf(w, "This is endpoint two with input: %s", input)
	})

	// Start the server and listen for requests on port 8000
	fmt.Println("Starting server on port 8000...")
	http.ListenAndServe(":8000", router)
}
