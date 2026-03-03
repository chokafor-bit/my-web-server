package main // Defines the main package (every executable Go program must have this)

import (
	"fmt"      // Package for formatted I/O (printing text)
	"net/http" // Package for building HTTP servers and handling web requests
)

// handlerFunc handles incoming HTTP requests
// w = response writer (used to send data back to the browser)
// res = the incoming request
func handlerFunc(w http.ResponseWriter, res *http.Request) {
	// Sends HTML content back to the client (browser)
	fmt.Fprint(w, "<h1> Welcome to my awesome site <h1>")
}

func main() {
	// Registers the handlerFunc to the "/" route (homepage)
	http.HandleFunc("/", handlerFunc)

	// Prints a message in the terminal to show the server is starting
	fmt.Println("starting the server on port :9000...")

	// Starts the web server on port 9000
	// nil means it uses the default router (http.DefaultServeMux)
	err := http.ListenAndServe(":9000", nil)

	// If there is an error starting the server, stop the program
	if err != nil {
		panic(err)
	}
}
