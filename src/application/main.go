package main

import (
	"fmt"
	"log"
	"net/http"
	"watervideodisplay/src/application/api"
	"watervideodisplay/src/application/cores"
	"watervideodisplay/src/application/cores/files"
)

func main() {
	cores.Init()
	log.Println("")
	log.Println("Water Display Start")
	core := cores.LoadEnv()
	files.SetResourceDir(core.GetResourceDir())
	// fmt.Println(core)

	http.HandleFunc("/", api.Index)
	http.HandleFunc("/content", api.Content)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", core.GetPort()), nil))
}
