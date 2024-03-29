package xfile

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
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
		result, err := Head(fn, 10)
		assert.Nil(t, err)
		assert.Equal(t, lines, result)
	}
	{
		result, err := Head(fn, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	}
	{
		result, err := Head(fn+"tom", 10)
		assert.True(t, os.IsNotExist(err))
		assert.Equal(t, 0, len(result))
	}
}

func TestHeadBytes(t *testing.T) {
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
		result, err := HeadBytes(fn, 30)
		assert.Nil(t, err)
		assert.Equal(t, []byte(strings.Join(lines, "\n")+"\n"), result)
	}
	{
		result, err := HeadBytes(fn, 6)
		assert.Nil(t, err)
		assert.Equal(t, []byte("line-1"), result)
	}
	{
		result, err := HeadBytes(fn, 7)
		assert.Nil(t, err)
		assert.Equal(t, []byte("line-1\n"), result)
	}
	{
		result, err := HeadBytes(fn, 0)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	}
	{
		result, err := HeadBytes(fn+"tom", 6)
		assert.True(t, os.IsNotExist(err))
		assert.Equal(t, 0, len(result))
	}
}
