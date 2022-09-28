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
	interval, err := strconv.ParseInt(os.Getenv("AWAKE_INTERVAL"), 0, 8)
	if err != nil {
		log.Fatalln("[ERR] Awake interval config", err)
	}
	core.SetAwakeInterval(interval)

	return core
}
