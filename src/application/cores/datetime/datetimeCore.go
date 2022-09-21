package datetime

import "time"

func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

func GetDateTodayAsString() string {
	return time.Now().Format("20060102")
}
