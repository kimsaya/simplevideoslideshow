package serial

import (
	"log"

	"github.com/tarm/serial"
)

var Message string = ""

func InitPort(port string) {
	go run(port)
}

func run(port string) {
	config := &serial.Config{
		Name:        port,
		Baud:        9600,
		ReadTimeout: 0,
		Size:        0,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal("[ERR] Connection com port error", port, err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		if err != nil {
			log.Println("[ERR] Read com port Error", err)
		}
		s := string(buf[:n])
		// fmt.Print(s)
		Message += s
	}
}
