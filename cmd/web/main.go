package main

import (
	"fmt"
	"net/http"

	"github.com/dmawardi/go-server/pkg/handlers"
)

const portNumber = ":8080"



func main() {
	// Route handler
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// The second argument is a handler, but not required due to above
	// returns an error
	
	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("error in serving port.")
	} 
} 