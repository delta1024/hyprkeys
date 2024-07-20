package main

// "locked": false,
//
//	"mouse": true,
//	"release": false,
//	"repeat": false,
//	"non_consuming": false,
//	"has_description": false,
//	"modmask": 64,
//	"submap": "",
//	"key": "mouse:273",
//	"keycode": 0,
//	"catch_all": false,
//	"description": "",
//	"dispatcher": "mouse",
//	"arg": "resizewindow"
type hyprKey struct {
	Locked         bool   `json:"locked"`
	Mouse          bool   `json:"mouse"`
	Release        bool   `json:"release"`
	Repeat         bool   `json:"repeat"`
	NonConsuming   bool   `json:"non_consuming"`
	HasDescription bool   `json:"has_description"`
	Modmask        int    `json:"modmask"`
	Submap         string `json:"submap"`
	Key            string `json:"key"`
	Keycode        int    `json:"keycode"`
	CatchAll       bool   `json:"catch_all"`
	Description    string `json:"description"`
	Dispatcher     string `json:"dispatcher"`
	Arg            string `json:"arg"`
}
