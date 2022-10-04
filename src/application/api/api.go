package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"watervideodisplay/src/application/assets"
	"watervideodisplay/src/application/cores/files"
	"watervideodisplay/src/application/services/player"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, files.GetWebIndexPath())
}

func Content(w http.ResponseWriter, r *http.Request) {
	content := new(assets.Content)
	content.Name = "Name"
	content.Url = "1"
	content.ItemType = "video"
	content.Duration = 100
	contentString, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(w, string(contentString))
}

func AppJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, files.GetAppJsPath())
}

func NextVideo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, player.GetPlayNext())
}

func PlayVideo(w http.ResponseWriter, r *http.Request) {
	value, _ := strconv.Atoi(r.URL.Query().Get("url"))
	http.ServeFile(w, r, player.GetPlay(value))
}
