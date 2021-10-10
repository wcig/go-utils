package xtime

import (
	"fmt"
	"testing"
)

func TestNowUnix(t *testing.T) {
	fmt.Println(NowUnix())
	fmt.Println(NowUnixMs())
	fmt.Println(NowUnixNs())
}

func TestNowDatetimeStr(t *testing.T) {
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
