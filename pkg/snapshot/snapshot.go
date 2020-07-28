package snapshot

import (
	"errors"

	"github.com/dg185200/qrappstore/pkg/app"
	"github.com/google/uuid"
)

// Snapshot is a pointer to an app with its invocation context specified
// this can be used to construct a URL for the QR code
type snapshot struct {
	ID  string   `json:"id,omitempty"`
	App *app.App `json:"app,omitempty"`
	URL string   `json:"url,omitempty"`

	// InvocationCtx a way to specify additional arguments for the instant app to load
	// these are passed in as query params to the url
	InvocationCtx map[string]string `json:"invocation_ctx,omitempty"`
}

// Cfg is an iterface to modify fields inside a snapshot
type Cfg interface {
	modify(*snapshot)
}

type snapshotOptFunc func(*snapshot)

func (f snapshotOptFunc) modify(s *snapshot) { f(s) }

func WithApp(app *app.App) Cfg {
	return snapshotOptFunc(func(s *snapshot) {
		s.App = app
	})
}

func WithURL(url string) Cfg {
	return snapshotOptFunc(func(s *snapshot) {
		s.URL = url
	})
}

func WithInvocationCtx(ictx map[string]string) Cfg {
	return snapshotOptFunc(func(s *snapshot) {
		s.InvocationCtx = ictx
	})
}

// NewWithOpts can be used to create a new snapshot of an App with the provided configurations
func NewWithOpts(cfgs ...Cfg) (*snapshot, error) {
	var snapshot snapshot
	for _, cfg := range cfgs {
		cfg.modify(&snapshot)
	}
	if snapshot.App == nil {
		return nil, errors.New("snapshot: app cannot be nil")
	}
	snapshot.ID = uuid.New().String()
	return &snapshot, nil
}
