package app

// Snapshot is a pointer to an app with its invocation context specified
// this can be used to construct a URL for the QR code
type Snapshot struct {
	App *App   `json:"app,omitempty"`
	URL string `json:"url,omitempty"`

	// InvocationCtx a way to specify additional arguments for the instant app to load
	// these are passed in as query params to the url
	InvocationCtx map[string]string `json:"invocation_ctx,omitempty"`
}
