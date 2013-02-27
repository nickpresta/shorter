package main

import (
	"flag"
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nickpresta/shorter/views"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

var (
	port            = flag.Int("port", 8080, "HTTP listen port")
	HTTPLogLocation = flag.String("log", "/tmp/shorter.log", "HTTP log file")
	DBConnection    = flag.String("connection", "localhost:28015", "RethinkDB connection (host:port)")
	HTTPLogger      io.Writer
)

func setupLogging(location string) {
	var err error
	HTTPLogger, err = os.OpenFile(location, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	setupLogging(*HTTPLogLocation)
	log.Printf("Now serving on http://localhost:%d\n", *port)

	runtime.GOMAXPROCS(runtime.NumCPU())

	session, err := r.Connect(*DBConnection, "urls")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	router := mux.NewRouter()
	// UUID format
	router.HandleFunc("/{key:[a-z0-9-]+}", func(w http.ResponseWriter, request *http.Request) {
		views.EmbiggenHandler(w, request, session)
	}).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		views.ShortenHandler(w, request, session)
	}).Methods("POST")
	router.HandleFunc("/", views.IndexHandler).Methods("GET")
	http.Handle("/", router)
	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.LoggingHandler(HTTPLogger, http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
