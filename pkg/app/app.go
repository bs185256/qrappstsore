package app

// App is a structure that represents the quick qr apps state
type App struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Cfg  []*Cfg `json:"cfg,omitempty"`
}

// Cfg is a way to provide **required** configuration for a snapshot
type Cfg struct {
	Name  string
	Value string
}
