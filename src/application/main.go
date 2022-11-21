package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"watervideodisplay/src/application/api"
	"watervideodisplay/src/application/assets"
	"watervideodisplay/src/application/cores"
	"watervideodisplay/src/application/cores/files"
	"watervideodisplay/src/application/cores/serial"
	"watervideodisplay/src/application/services/player"
)

var core *assets.Core

func main() {
	cores.Init()
	log.Println("")
	log.Println("Water Display Start")
	core = cores.LoadEnv()
	files.SetResourceDir(core.GetResourceDir())
	player.IsAlwaysOn = core.IsAlwaysOn()
	// fmt.Println(core)
	serial.InitPort(core.GetComPort())
	go awake()

	http.HandleFunc("/", api.Index)
	http.HandleFunc("/script/app", api.AppJs)
	http.HandleFunc("/content", api.Content)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", core.GetPort()), nil))

	// This class can build only in Linux env
	// opio.InitInterruptor()
}

func awake() {
	for {
		time.Sleep(time.Duration(1 * int64(time.Second)))
		if len(serial.Message) > 0 {
			serial.Message = ""
			player.IsTimeOff = false
			time.Sleep(time.Duration(core.GetAwakeInterval() * int64(time.Second)))
		} else {
			player.IsTimeOff = true
		}
	}
}
