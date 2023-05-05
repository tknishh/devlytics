package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// endpoint-one
	router.HandleFunc("/endpoint-one", func(w http.ResponseWriter, r *http.Request) {
		// response to client
		fmt.Fprint(w, "This is endpoint one!")
	})

	// endpoint-two
	router.HandleFunc("/endpoint-two/{input}", func(w http.ResponseWriter, r *http.Request) {
		// get input from url
		vars := mux.Vars(r)
		input := vars["input"]

		// response to client
		fmt.Fprintf(w, "This is endpoint two with input: %s", input)
	})

	// Start the server and listen for requests on port 8000
	fmt.Println("Starting server on port 8000...")
	http.ListenAndServe(":8000", router)
}
