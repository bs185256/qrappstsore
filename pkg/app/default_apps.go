package app

// Default is a map of defaults apps offered by the QR app store
var Default = map[string]*App{
	"quick-order": {
		ID:   "quick-order",
		Name: "NCR quick order",
		Cfg: []*Cfg{
			{"headerColor", "#3B5998"},
			{"footerColor", "#3B5998"},
		},
	},
	"easy-shopper": {
		ID:   "easy-shopper",
		Name: "Easy Shopper",
		Cfg: []*Cfg{
			{"headerColor", "#3B5998"},
			{"footerColor", "#3B5998"},
		},
	},
}
