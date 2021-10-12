package xjson

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	v := map[string]int{"id": 1}
	fmt.Println(ToString(v))
	fmt.Println(ToPrettyString(v))
}

func TestPrintString(t *testing.T) {
	v := map[string]int{"id": 1}
	PrintString(v)
	PrintPrettyString(v)
}
