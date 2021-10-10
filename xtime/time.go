package xtime

import "time"

// now time timestamp (millisecond)
func GetNowMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// now time timestamp (nanosecond)
func GetNowNano() int64 {
	return time.Now().UnixNano()
}

// now string (default YYYY-MM-dd HH:mm:ss)
func GetNowStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// today begin time
func GetTodayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// today end time
func GetTodayEndTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, now.Location())
}
