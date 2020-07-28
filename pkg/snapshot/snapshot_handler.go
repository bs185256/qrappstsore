package snapshot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dg185200/qrappstore/pkg/app"
	"github.com/gorilla/mux"
)

type addSnapshotHandler struct {
	library Library
}

func NewAddSnapshotHandler(library Library) *addSnapshotHandler {
	return &addSnapshotHandler{library: library}
}

func (h *addSnapshotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		App *app.App          `json:"app,omitempty"`
		URL string            `json:"url,omitempty"`
		Ctx map[string]string `json:"ctx,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		log.Println(err)
		return
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
}

type getSnapshotsHandler struct {
	library Library
}

func NewGetSnapshotsHandler(library Library) *getSnapshotsHandler {
	return &getSnapshotsHandler{library: library}
}

func (h *getSnapshotsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	s, _ := h.library.Get(id)
	json.NewEncoder(w).Encode(s)
}
