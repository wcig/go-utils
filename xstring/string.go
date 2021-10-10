package xstring

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrConvert = errors.New("convert string err")
)

// str is empty
func IsEmpty(s string) bool {
	return s == ""
}

// str is not empty
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// str is blank
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// str is not blank
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// string -> int
func GetInt(s string, defaultValue int) int {
	if v, err := strToType(s, reflect.Int); v != nil && err == nil {
		if value, ok := v.(int); ok {
			return value
		}
	}
	return defaultValue
}

// string -> int64
func GetInt64(s string, defaultValue int64) int64 {
	if v, err := strToType(s, reflect.Int64); v != nil && err == nil {
		if value, ok := v.(int64); ok {
			return value
		}
	}
	return defaultValue
}

// string -> float32
func GetFloat32(s string, defaultValue float32) float32 {
	if v, err := strToType(s, reflect.Float32); v != nil && err == nil {
		if value, ok := v.(float32); ok {
			return value
		}
	}
	return defaultValue
}

// string -> float64
func GetFloat64(s string, defaultValue float64) float64 {
	if v, err := strToType(s, reflect.Float64); v != nil && err == nil {
		if value, ok := v.(float64); ok {
			return value
		}
	}
	return defaultValue
}

// str to type val
func strToType(s string, kind reflect.Kind) (interface{}, error) {
	var err error
	if len(s) <= 0 {
		return nil, ErrConvert
	}

	switch kind {
	case reflect.Int:
		var v int
		v, err = strconv.Atoi(s)
		return v, err
	case reflect.Int64:
		var v int64
		v, err = strconv.ParseInt(s, 10, 64)
		return v, err
	case reflect.Uint:
		var v uint64
		v, err = strconv.ParseUint(s, 10, 64)
		return v, err
	case reflect.Float32:
		var v float64
		v, err = strconv.ParseFloat(s, 32)
		return float32(v), err
	case reflect.Float64:
		var v float64
		v, err = strconv.ParseFloat(s, 64)
		return v, err
	}
	return nil, ErrConvert
}

// int -> str
func IntToStr(n int) string {
	return strconv.Itoa(n)
}

// int64 -> str
func Int64ToStr(n int64) string {
	return strconv.FormatInt(n, 10)
}
