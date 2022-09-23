package opio

// This class can build only in Linux env
import (
	"time"
	"watervideodisplay/src/application/cores/datetime"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var IsInterrupted bool = false

var EndDuration int64 = 5000
var interruptStartTime int64 = 0

func InitInterruptor() {
	r := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(r, "11")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			interruptStartTime = datetime.GetNowTimestamp()
			IsInterrupted = true
		})

		// button.On(gpio.ButtonRelease, func(data interface{}) {
		// 	fmt.Println("button released")
		// 	led.Off()
		// })
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
		work,
	)

	robot.Start()
	go run()
}

func run() {
	for {
		time.Sleep(time.Duration(EndDuration))
		interruptTimer()
	}
}

func OnInterrupted() {
	IsInterrupted = true
}

func interruptTimer() {
	if !IsInterrupted {
		return
	}
	if datetime.GetNowTimestamp()-interruptStartTime > EndDuration {
		IsInterrupted = false
	}
}
