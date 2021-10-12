package xtime

import (
	"fmt"
	"testing"
	"time"
)

func TestNowUnix(t *testing.T) {
	fmt.Println(NowUnix())
	fmt.Println(NowUnixMs())
	fmt.Println(NowUnixNs())
}

func TestNowDatetimeStr(t *testing.T) {
	fmt.Println(NowDateStr())
	fmt.Println(NowDatetimeStr())
	fmt.Println(NowDatetimeMsStr())
	fmt.Println(NowDatetimeNsStr())
}

func TestBeginEnd(t *testing.T) {
	fmt.Println(BeginOfYear())
	fmt.Println(EndOfYear())

	fmt.Println(BeginOfMonth())
	fmt.Println(EndOfMonth())

	fmt.Println(BeginOfWeek())
	fmt.Println(EndOfWeek())

	fmt.Println(BeginOfDay())
	fmt.Println(EndOfDay())

	fmt.Println(BeginOfHour())
	fmt.Println(EndOfHour())

	fmt.Println(BeginOfMinute())
	fmt.Println(EndOfMinute())
}

func TestUnixToTime(t *testing.T) {
	now := time.Now()
	sec := now.Unix()
	nsec := now.UnixNano()
	msec := nsec / 1e6

	fmt.Println(UnixToTime(sec))
	fmt.Println(UnixMsToTime(msec))
	fmt.Println(UnixNsToTime(nsec))
}

func TestTimeToUnix(t *testing.T) {
	now := time.Now()
	fmt.Println(TimeToUnix(now))
	fmt.Println(TimeToUnixMs(now))
	fmt.Println(TimeToUnixNs(now))
}

func TestDatetimeStrToTime(t *testing.T) {
	now := time.Now()
	v1 := TimeToDateStr(now)
	v2 := TimeToDatetimeStr(now)
	v3 := TimeToDatetimeMsStr(now)
	v4 := TimeToDatetimeNsStr(now)

	fmt.Println(DateStrToTime(v1))
	fmt.Println(DatetimeStrToTime(v2))
	fmt.Println(DatetimeMsStrToTime(v3))
	fmt.Println(DatetimeNsStrToTime(v4))
}

func TestDateToDatetimeStr(t *testing.T) {
	now := time.Now()
	fmt.Println(TimeToDateStr(now))
	fmt.Println(TimeToDatetimeStr(now))
	fmt.Println(TimeToDatetimeMsStr(now))
	fmt.Println(TimeToDatetimeNsStr(now))
}

func TestLoc(t *testing.T) {
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().In(UTCLoc).Location())
	fmt.Println(time.Now().In(BeijingLoc).Location())

	fmt.Println(time.Now().Format(DatetimeFormat))
	fmt.Println(time.Now().In(UTCLoc).Format(DatetimeFormat))
	fmt.Println(time.Now().In(BeijingLoc).Format(DatetimeFormat))
	// output:
	// Local
	// UTC
	// UTC+8
	// 2021-10-12 22:59:58
	// 2021-10-12 14:59:58
	// 2021-10-12 22:59:58
}
