package snapshot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dg185200/qrappstore/internal/httperror"
	"github.com/dg185200/qrappstore/pkg/app"
	"github.com/gorilla/mux"
)

type addSnapshotHandler struct {
	library Library
}

func NewAddSnapshotHandler(library Library) *addSnapshotHandler {
	return &addSnapshotHandler{library: library}
}

func (h *addSnapshotHandler) HandleRequest(w http.ResponseWriter, r *http.Request) error {
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
	ss, err := NewWithOpts(WithApp(reqData.App), WithURL(reqData.URL), WithInvocationCtx(reqData.Ctx))
	if err != nil {
		log.Println(err)
	}

	ss, err = h.library.Add(ss)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(ss)
	return nil
}

type getSnapshotsHandler struct {
	library Library
}

func NewGetSnapshotsHandler(library Library) *getSnapshotsHandler {
	return &getSnapshotsHandler{library: library}
}

func (h *getSnapshotsHandler) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	s, err := h.library.Get(id)
	if err != nil {
		return httperror.StatusError{Code: 404, Err: err}
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		return httperror.StatusError{Code: 500, Err: err}
	}
	return nil
}
