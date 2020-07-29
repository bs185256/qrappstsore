package snapshot

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dg185200/qrappstore/internal/httperror"
	"github.com/dg185200/qrappstore/pkg/app"
	"github.com/gorilla/mux"
)

const (
	nepOrganization = "nep-organization"
)

type addSnapshotHandler struct {
	library Library
}

func NewAddSnapshotHandler(library Library) *addSnapshotHandler {
	return &addSnapshotHandler{library: library}
}

func (h *addSnapshotHandler) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	organization := r.Header.Get(nepOrganization)
	if organization == "" {
		return httperror.StatusError{Code: 400, Err: errors.New("snapshot: header nep-organization is required")}
	}
	var reqData struct {
		App *app.App          `json:"app,omitempty"`
		URL string            `json:"url,omitempty"`
		Ctx map[string]string `json:"ctx,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		log.Println(err)
		return httperror.StatusError{Code: 500, Err: err}
	}
	defer r.Body.Close()
	snapshot, err := NewWithOpts(WithApp(reqData.App), WithURL(reqData.URL),
		WithInvocationCtx(reqData.Ctx), withOrganization((organization)))
	if err != nil {
		log.Println(err)
	}

	snapshot, err = h.library.Add(snapshot)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(snapshot)
	return nil
}

type getSnapshotsHandler struct {
	library Library
}

func NewGetSnapshotsHandler(library Library) *getSnapshotsHandler {
	return &getSnapshotsHandler{library: library}
}

func (h *getSnapshotsHandler) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	organization := r.Header.Get(nepOrganization)
	if organization == "" {
		return httperror.StatusError{Code: 400, Err: errors.New("snapshot: header nep-organization is required")}
	}
	vars := mux.Vars(r)
	id := vars["id"]
	key := fmt.Sprintf("%s-%s", organization, id)
	s, err := h.library.Get(key)
	if err != nil {
		return httperror.StatusError{Code: 404, Err: err}
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		return httperror.StatusError{Code: 500, Err: err}
	}
	return nil
}
