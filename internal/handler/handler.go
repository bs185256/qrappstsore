package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dg185200/qrappstore/internal/httperror"
)

type RequestHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request) error
}

func New(h RequestHandler) http.Handler {
	return handler{H: h}
}

type handler struct {
	H RequestHandler
}

type errResponse struct {
	Status int    `json:"status,omitempty"`
	Msg    string `json:"error,omitempty"`
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.H.HandleRequest(w, r); err != nil {
		switch e := err.(type) {
		case httperror.Error:
			log.Println("error while fulfilling request:", err)
			resp := &errResponse{
				Status: e.Status(),
				Msg:    e.Error(),
			}
			w.WriteHeader(e.Status())
			json.NewEncoder(w).Encode(resp)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
