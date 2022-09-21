package gpio

import "watervideodisplay/src/application/cores/datetime"

var IsInterrupted bool = false

var EndDuration int64 = 5000
var interruptStartTime int64 = 0

func InitInterruptor() {
	go run()
}

func run() {

}

func OnInterrupted() {
	IsInterrupted = true
}

func interruptTimer() bool {
	if datetime.GetNowTimestamp()-interruptStartTime > EndDuration {
		IsInterrupted = false
	}
	return !IsInterrupted
}
