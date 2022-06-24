package xfile

import (
	"os"
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

	// check
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
