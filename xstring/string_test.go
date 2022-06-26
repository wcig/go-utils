package xstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteCount(t *testing.T) {
	var ts = []struct {
		input  string
		expect int
	}{
		{
			input:  "12345",
			expect: 5,
		},
		{
			input:  "你好tom",
			expect: 9,
		},
	}
	for _, v := range ts {
		count := ByteCount(v.input)
		assert.Equal(t, v.expect, count)
	}
}

func TestRuneCount(t *testing.T) {
	var ts = []struct {
		input  string
		expect int
	}{
		{
			input:  "12345",
			expect: 5,
		},
		{
			input:  "你好tom",
			expect: 5,
		},
	}
	for _, v := range ts {
		count := RuneCount(v.input)
		assert.Equal(t, v.expect, count)
	}
}

func TestRuneSub(t *testing.T) {
	var ts = []struct {
		input  string
		begin  int
		length int
		expect string
	}{
		{
			input:  "12345",
			begin:  0,
			length: 5,
			expect: "12345",
		},
		{
			input:  "12345",
			begin:  0,
			length: 10,
			expect: "12345",
		},
		{
			input:  "12345",
			begin:  2,
			length: 3,
			expect: "345",
		},
		{
			input:  "12345",
			begin:  2,
			length: 10,
			expect: "345",
		},
		{
			input:  "12345",
			begin:  -1,
			length: 5,
			expect: "",
		},

		{
			input:  "你好tom",
			begin:  0,
			length: 5,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			begin:  0,
			length: 10,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			begin:  1,
			length: 3,
			expect: "好to",
		},
		{
			input:  "你好tom",
			begin:  1,
			length: 10,
			expect: "好tom",
		},
		{
			input:  "你好tom",
			begin:  -1,
			length: 5,
			expect: "",
		},
	}
	for _, v := range ts {
		val := RuneSub(v.input, v.begin, v.length)
		assert.Equal(t, v.expect, val)
	}
}

func TestRuneSubLeft(t *testing.T) {
	var ts = []struct {
		input  string
		length int
		expect string
	}{
		{
			input:  "12345",
			length: 3,
			expect: "123",
		},
		{
			input:  "12345",
			length: 5,
			expect: "12345",
		},
		{
			input:  "12345",
			length: 10,
			expect: "12345",
		},
		{
			input:  "12345",
			length: 0,
			expect: "",
		},
		{
			input:  "12345",
			length: -1,
			expect: "",
		},
		{
			input:  "你好tom",
			length: 3,
			expect: "你好t",
		},
		{
			input:  "你好tom",
			length: 5,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			length: 10,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			length: 0,
			expect: "",
		},
		{
			input:  "你好tom",
			length: -1,
			expect: "",
		},
	}
	for _, v := range ts {
		val := RuneSubLeft(v.input, v.length)
		assert.Equal(t, v.expect, val)
	}
}

func TestRuneSubRight(t *testing.T) {
	var ts = []struct {
		input  string
		length int
		expect string
	}{
		{
			input:  "12345",
			length: 3,
			expect: "345",
		},
		{
			input:  "12345",
			length: 5,
			expect: "12345",
		},
		{
			input:  "12345",
			length: 10,
			expect: "12345",
		},
		{
			input:  "12345",
			length: 0,
			expect: "",
		},
		{
			input:  "12345",
			length: -1,
			expect: "",
		},
		{
			input:  "你好tom",
			length: 4,
			expect: "好tom",
		},
		{
			input:  "你好tom",
			length: 5,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			length: 10,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			length: 0,
			expect: "",
		},
		{
			input:  "你好tom",
			length: -1,
			expect: "",
		},
	}
	for _, v := range ts {
		val := RuneSubRight(v.input, v.length)
		assert.Equal(t, v.expect, val)
	}
}

func TestRuneSubString(t *testing.T) {
	var ts = []struct {
		input  string
		begin  int
		end    int
		expect string
	}{
		{
			input:  "12345",
			begin:  0,
			end:    5,
			expect: "12345",
		},
		{
			input:  "12345",
			begin:  0,
			end:    10,
			expect: "12345",
		},
		{
			input:  "12345",
			begin:  2,
			end:    4,
			expect: "34",
		},
		{
			input:  "12345",
			begin:  -1,
			end:    5,
			expect: "",
		},
		{
			input:  "12345",
			begin:  2,
			end:    1,
			expect: "",
		},
		{
			input:  "你好tom",
			begin:  0,
			end:    5,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			begin:  0,
			end:    10,
			expect: "你好tom",
		},
		{
			input:  "你好tom",
			begin:  2,
			end:    4,
			expect: "to",
		},
		{
			input:  "你好tom",
			begin:  -1,
			end:    5,
			expect: "",
		},
		{
			input:  "你好tom",
			begin:  2,
			end:    1,
			expect: "",
		},
	}
	for _, v := range ts {
		val := RuneSubString(v.input, v.begin, v.end)
		assert.Equal(t, v.expect, val)
	}
}
