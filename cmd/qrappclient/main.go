package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:9001", http.FileServer(http.Dir("web/businessOwnerLogin"))))
}
