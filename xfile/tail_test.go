package xfile

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTail(t *testing.T) {
	// prepare
	fn := "tmp.txt"
	f, err := os.Create(fn)
	if err != nil {
		t.Fatal("open filed:", err)
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(fn)
	}()

	lines := []string{"line-1", "line-2", "line-3"}
	for _, line := range lines {
		if _, err = f.WriteString(line + "\n"); err != nil {
			t.Fatal("write failed:", err)
		}
	}

	// test
	{
		result, err := Tail(fn, 1)
		assert.Nil(t, err)
		assert.Equal(t, []string{lines[2]}, result)
	}
	{
		result, err := Tail(fn, 3)
		assert.Nil(t, err)
		assert.Equal(t, lines, result)
	}
	{
		result, err := Tail(fn, 10)
		assert.Nil(t, err)
		assert.Equal(t, lines, result)
	}
	{
		result, err := Tail(fn, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	}
	{
		result, err := Tail(fn+"tom", 10)
		assert.True(t, os.IsNotExist(err))
		assert.Equal(t, 0, len(result))
	}
}

func TestTailBytes(t *testing.T) {
	// prepare
	fn := "tmp.txt"
	f, err := os.Create(fn)
	if err != nil {
		t.Fatal("open filed:", err)
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(fn)
	}()

	lines := []string{"line-1", "line-2", "line-3"}
	for _, line := range lines {
		if _, err = f.WriteString(line + "\n"); err != nil {
			t.Fatal("write failed:", err)
		}
	}

	// test
	{
		result, err := TailBytes(fn, 30)
		assert.Nil(t, err)
		assert.Equal(t, []byte(strings.Join(lines, "\n")), result)
	}
	{
		result, err := TailBytes(fn, 6)
		assert.Nil(t, err)
		assert.Equal(t, []byte("line-3"), result)
	}
	{
		result, err := TailBytes(fn, 7)
		assert.Nil(t, err)
		assert.Equal(t, []byte("\nline-3"), result)
	}
	{
		result, err := TailBytes(fn, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	}
	{
		result, err := TailBytes(fn+"tom", 6)
		assert.True(t, os.IsNotExist(err))
		assert.Equal(t, 0, len(result))
	}
}
