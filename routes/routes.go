package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "timechip API")
	fmt.Fprintf(os.Stdout, "timechip API stdout\n")
	fmt.Println("ahoj")
}

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")

	router.HandleFunc("/homepage/nejblizsi-zavody/{race-year}", Neco).Methods("GET")
	router.HandleFunc("/homepage/zavody/{race-year}", Zavody).Methods("GET")

	staticFileDirectory := http.Dir("./static/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return router
}
