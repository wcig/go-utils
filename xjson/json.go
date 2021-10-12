package xjson

import (
	"encoding/json"
	"fmt"
)

// ToString convert v to json string
func ToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// ToPrettyString convert v to pretty json string
func ToPrettyString(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

// PrintString print v json string
func PrintString(v interface{}) {
	fmt.Println(ToString(v))
}

// PrintPrettyString print v pretty json string
func PrintPrettyString(v interface{}) {
	fmt.Println(ToPrettyString(v))
}
