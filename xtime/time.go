package xtime

import "time"

const (
	datetimeFormat   = "2006-01-02 15:04:05"
	datetimeMsFormat = "2006-01-02 15:04:05.999"
	datetimeNsFormat = "2006-01-02 15:04:05.999999999"
)

// NowUnix now timestamp (second)
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowUnixMs now timestamp (millisecond)
func NowUnixMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// NowUnixNs now timestamp (nanosecond)
func NowUnixNs() int64 {
	return time.Now().UnixNano()
}

// NowDatetimeStr now datetime string
func NowDatetimeStr() string {
	return time.Now().Format(datetimeFormat)
}

// NowDatetimeMsStr now datetime string (millisecond)
func NowDatetimeMsStr() string {
	return time.Now().Format(datetimeMsFormat)
}

// NowDatetimeNsStr now datetime string (nanosecond)
func NowDatetimeNsStr() string {
	return time.Now().Format(datetimeNsFormat)
}

// BeginOfYear begin of year
func BeginOfYear() time.Time {
	now := time.Now()
	return time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
}

// EndOfYear end of year
func EndOfYear() time.Time {
	return BeginOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// BeginOfMonth begin of month
func BeginOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// EndOfMonth end of month
func EndOfMonth() time.Time {
	return BeginOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfWeek begin of week (week begin from Sunday)
func BeginOfWeek() time.Time {
	now := time.Now()
	y, m, d := now.AddDate(0, 0, 0-int(now.Weekday())).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}

// EndOfWeek end of week (week begin from Sunday)
func EndOfWeek() time.Time {
	return BeginOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// BeginOfDay begin of day
func BeginOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// EndOfDay end of day
func EndOfDay() time.Time {
	return BeginOfDay().AddDate(0, 0, 1).Add(-time.Nanosecond)
}

// BeginOfHour begin of hour
func BeginOfHour() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
}

// EndOfHour end of hour
func EndOfHour() time.Time {
	return BeginOfHour().Add(time.Hour).Add(-time.Nanosecond)
}

// BeginOfMinute begin of minute
func BeginOfMinute() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
}

// EndOfMinute end of minute
func EndOfMinute() time.Time {
	return BeginOfMinute().Add(time.Minute).Add(-time.Nanosecond)
}
