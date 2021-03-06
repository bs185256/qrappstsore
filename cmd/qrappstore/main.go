package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/dg185200/qrappstore/internal/handler"
	"github.com/dg185200/qrappstore/pkg/app"
	"github.com/dg185200/qrappstore/pkg/items"
	"github.com/dg185200/qrappstore/pkg/snapshot"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	// host+port
	addr     string
	debug    bool
	username string
	password string
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:9000", "host+port for the server to bind to")
	flag.BoolVar(&debug, "debug", false, "enable this flag for all debug logging")
	flag.StringVar(&username, "username", "", "nep username")
	flag.StringVar(&password, "password", "", "nep password")
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// set up dependencies
	libray := snapshot.NewLibrary()

	r := mux.NewRouter()
	r.Use(timer, requestLoggingMiddleWare)
	r.HandleFunc("/_ah/healthz", healthHandler).Methods(http.MethodGet)

	sh := http.StripPrefix("/apps", http.FileServer(http.Dir("web/businessOwnerLogin")))
	r.PathPrefix("/apps").Handler(sh)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(contentTypeJSON)
	apiRouter.HandleFunc("/apps", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(app.Default); err != nil {
			log.Println(err)
		}
	}).Methods(http.MethodGet)

	apiRouter.Handle("/snapshots", handler.New(snapshot.NewAddSnapshotHandler(libray))).Methods(http.MethodPost)
	apiRouter.Handle("/snapshots/{id}", handler.New(snapshot.NewGetSnapshotsHandler(libray))).Methods(http.MethodGet)
	apiRouter.Handle("/items", handler.New(items.NewHandler(username, password))).Methods(http.MethodGet)

	log.Println("starting server on:", addr)
	log.Fatal(http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, r)))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func contentTypeJSON(nh http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		nh.ServeHTTP(w, r)
	})
}

func timer(nh http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		nh.ServeHTTP(w, r)
		log.Println("took", time.Since(start), "to fullfill request", r.URL.EscapedPath())
	})
}

func requestLoggingMiddleWare(nh http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if debug {
			b, _ := httputil.DumpRequest(r, true)
			log.Printf("\n----------\n%s\n----------\n", b)
		}
		// move on the the next handler
		nh.ServeHTTP(w, r)
	})
}
