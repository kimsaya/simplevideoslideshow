package api

import (
	"log"
	"net/http"
	"watervideodisplay/src/application/cores/files"
	"watervideodisplay/src/application/services/player"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fileBytes, err := files.ReadFile(files.GetWebIndexPath())
	if err != nil {
		log.Println("[ERR] Read : ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	// return
}

func Content(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fileBytes, err := files.ReadFile(player.GetPlayNext())
	if err != nil {
		log.Println("[ERR] Read : ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	// return
}
