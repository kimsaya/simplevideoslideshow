package cores

import (
	"log"
	"os"
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
	core.SetResourceDir(os.Getenv("RESOURCE_DIR"))

	return core
}
