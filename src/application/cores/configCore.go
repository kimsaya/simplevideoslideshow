package cores

import (
	"log"
	"os"
	"strconv"
	"watervideodisplay/src/application/assets"
	"watervideodisplay/src/application/cores/files"

	"github.com/joho/godotenv"
)

func LoadEnv() *assets.Core {
	path := files.GetEnvPath()
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalln("[ERR] Loading .env file: " + path)
	}

	core := new(assets.Core)
	core.SetPort(os.Getenv("PORT"))
	core.SetComPort(os.Getenv("COM_PORT"))
	core.SetResourceDir(os.Getenv("RESOURCE_DIR"))
	interval, err := strconv.ParseInt(os.Getenv("AWAKE_INTERVAL"), 0, 32)
	if err != nil {
		interval = 1
	}
	core.SetAwakeInterval(interval)
	isAlwaysOn, err := strconv.ParseBool(os.Getenv("ALWAYS_ON"))
	if err != nil {
		isAlwaysOn = false
	}
	core.SetAlwaysOn(isAlwaysOn)

	return core
}
