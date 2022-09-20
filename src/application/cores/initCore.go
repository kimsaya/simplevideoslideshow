package cores

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"watervideodisplay/src/application/cores/datetime"
	"watervideodisplay/src/application/cores/files"
)

func Init() {
	initLogEnv()
	initConfigEnv()
	initWebEnv()
	initResourceEnv()
}

func initLogEnv() {
	logDir := files.GetCurrentDirectory() + "logs/"
	logFile := logDir + datetime.GetDateTodayAsString() + ".md"
	files.CreateDirectory(logDir)
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("[ERR] Opening logfile: ", err)
	}
	log.SetOutput(f)
}

func initConfigEnv() {
	configDir := files.GetCurrentDirectory() + "configs/"
	files.CreateDirectory(configDir)
	configFile := configDir + "application.env"
	if !files.IsFileExit(configFile) {
		// files.CreateFile(configFile)
		files.WriteFile(configFile, "PORT=8080")
	}
}

func initWebEnv() {
	webDir := files.GetCurrentDirectory() + "webs"
	indexFile := files.GetWebIndexPath()
	videoEndFile := files.GetVideoEndedPath()
	files.CreateDirectory(webDir)
	isErr := false
	if !files.IsFileExit(indexFile) {
		errorMissingComponent("index.html")
	}
	if !files.IsFileExit(videoEndFile) {
		errorMissingComponent("end.mp4")
	}
	if isErr {
		log.Fatalln("Web Env missing component")
		bufio.NewScanner(os.Stdin)
		os.Exit(0)
	}
}

func initResourceEnv() {
	resourceDir := files.GetResourceDirectory()
	files.CreateDirectory(resourceDir)
}

func errorMissingComponent(value string) {
	message := "Enviroment missing component!: " + value
	log.Println(message)
	fmt.Println(message)
}
