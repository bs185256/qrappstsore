package snapshot

import (
	"errors"
	"fmt"

	"github.com/dg185200/qrappstore/pkg/app"
	"github.com/google/uuid"
)

// Snapshot is a pointer to an app with its invocation context specified
// this can be used to construct a URL for the QR code
type snapshot struct {
	ID           string `json:"id,omitempty"`
	key          string
	organization string
	App          *app.App `json:"app,omitempty"`
	URL          string   `json:"url,omitempty"`

	// InvocationCtx a way to specify additional arguments for the instant app to load
	// these are passed in as query params to the url
	InvocationCtx map[string]string `json:"invocation_ctx,omitempty"`
}

// Cfg is an iterface to modify fields inside a snapshot
type Cfg interface {
	modify(*snapshot)
}

type snapshotCfgFunc func(*snapshot)

func (f snapshotCfgFunc) modify(s *snapshot) { f(s) }

func WithApp(app *app.App) Cfg {
	return snapshotCfgFunc(func(s *snapshot) {
		s.App = app
	})
}

func WithURL(url string) Cfg {
	return snapshotCfgFunc(func(s *snapshot) {
		s.URL = url
	})
}

func WithInvocationCtx(ictx map[string]string) Cfg {
	return snapshotCfgFunc(func(s *snapshot) {
		s.InvocationCtx = ictx
	})
}

func withOrganization(org string) Cfg {
	return snapshotCfgFunc(func(s *snapshot) {
		s.organization = org
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
	if snapshot.organization == "" {
		return nil, errors.New("snapshot: organization cannot be nil while building snapshot")
	}
	snapshot.ID = uuid.New().String()
	snapshot.key = fmt.Sprintf("%s-%s", snapshot.organization, snapshot.ID)
	return &snapshot, nil
}
