package datetime

import "time"

func GetDateTodayAsString() string {
	return time.Now().Format("20060102")
}
