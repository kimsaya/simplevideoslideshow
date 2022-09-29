package api

import (
	"net/http"
	"watervideodisplay/src/application/cores/files"
	"watervideodisplay/src/application/services/player"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, files.GetWebIndexPath())
}

func Content(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, player.GetPlayNext())
}

func AppJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, files.GetAppJsPath())
}
