package xfile

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testTailLines = [][]string{
		{"line-1\n", "line-2\n", "line-3\n"},
		{"line-1\n", "line-2\n", "line-3"},
		{"line-1\n", "line-2\n", "\n"},
		{"line-1\n", "\n", "\n"},
		{"\n", "\n", "\n"},
	}

	testTailBytes = [][]byte{
		[]byte("line1\nline2\nline3\n"),
		[]byte("line1\nline2\nline3"),
		[]byte("line1\n\n\n"),
		[]byte("\n\n\n"),
	}
)

func TestTail(t *testing.T) {
	fn := "tmp.txt"
	defer func() {
		_ = os.Remove(fn)
	}()

	for _, lines := range testTailLines {
		err := ioutil.WriteFile(fn, []byte(strings.Join(lines, "")), 0644)
		assert.Nil(t, err)

		result, err := Tail(fn, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))

		result, err = Tail(fn, 1)
		assert.Nil(t, err)
		assert.Equal(t, []string{lines[2]}, result)

		result, err = Tail(fn, 2)
		assert.Nil(t, err)
		assert.Equal(t, []string{lines[1], lines[2]}, result)

		result, err = Tail(fn, 3)
		assert.Nil(t, err)
		assert.Equal(t, lines, result)

		result, err = Tail(fn, 10)
		assert.Nil(t, err)
		assert.Equal(t, lines, result)

		result, err = Tail(fn+"tom", 10)
		assert.True(t, os.IsNotExist(err))
		assert.Equal(t, 0, len(result))
	}
}

func TestTailV2(t *testing.T) {
	fn := "tmp.txt"
	defer func() {
		_ = os.Remove(fn)
	}()

	for _, lines := range testTailLines {
		err := ioutil.WriteFile(fn, []byte(strings.Join(lines, "")), 0644)
		assert.Nil(t, err)

		result, err := TailV2(fn, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))

		result, err = TailV2(fn, 1)
		assert.Nil(t, err)
		assert.Equal(t, []string{lines[2]}, result)

		result, err = TailV2(fn, 2)
		assert.Nil(t, err)
		assert.Equal(t, []string{lines[1], lines[2]}, result)

		result, err = TailV2(fn, 3)
		assert.Nil(t, err)
		assert.Equal(t, lines, result)

		result, err = TailV2(fn, 10)
		assert.Nil(t, err)
		assert.Equal(t, lines, result)

		result, err = Tail(fn+"tom", 10)
		assert.True(t, os.IsNotExist(err))
		assert.Equal(t, 0, len(result))
	}
}

func TestTailBytes(t *testing.T) {
	fn := "tmp.txt"
	defer func() {
		_ = os.Remove(fn)
	}()

	for _, bytes := range testTailBytes {
		err := ioutil.WriteFile(fn, bytes, 0644)
		assert.Nil(t, err)

		for i := 1; i < len(bytes); i++ {
			result, err := TailBytes(fn, i)
			assert.Nil(t, err)
			assert.Equal(t, bytes[len(bytes)-i:], result)
		}

		result, err := TailBytes(fn, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))

		result, err = TailBytes(fn, 100)
		assert.Nil(t, err)
		assert.Equal(t, bytes, result)
	}
}
