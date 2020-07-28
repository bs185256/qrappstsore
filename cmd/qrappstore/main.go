package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

var (
	// host+port
	addr  string
	debug bool
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:9000", "host+port for the server to bind to")
	flag.BoolVar(&debug, "debug", true, "enable this flag for all debug logging")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()
	r.Use(requestLoggingMiddleWare)
	r.HandleFunc("/_ah/healthz", healthHandler).Methods("GET")

	log.Println("starting server on:", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
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
