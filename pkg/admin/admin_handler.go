package admin

import (
	"log"
	"net/http"
	"os"
)

type adminHandler struct {
	root string
}

func Handler(root string) http.Handler {
	return &adminHandler{root}
}

func (h adminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(os.Getwd())
	http.FileServer(http.Dir("web/"+h.root)).ServeHTTP(w, r)
}
