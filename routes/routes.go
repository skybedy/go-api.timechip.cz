package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"go-api.timechip.cz/conf"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "timechip API")
	fmt.Fprintf(os.Stdout, "timechip API stdout\n")
	fmt.Println("ahoj")
}

func newLoggingHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(dst, h)
	}
}

func NewRouter() *mux.Router {
	logFile, err := os.OpenFile(conf.AppPath+"/log/server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal()
	}
	loggingHandler := newLoggingHandler(logFile)
	stdChain := alice.New(loggingHandler)
	router := mux.NewRouter()
	router.Handle("/", stdChain.Then(http.HandlerFunc(Index)))
	router.Handle("/homepage/nejblizsi-zavody/{race-year}", stdChain.Then(http.HandlerFunc(Neco))).Methods("GET")
	router.Handle("/homepage/zavody/{race-year}", stdChain.Then(http.HandlerFunc(Zavody))).Methods("GET")

	staticFileDirectory := http.Dir(conf.AppPath + "/static/")
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
