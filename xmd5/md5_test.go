package xmd5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcig/go-utils/xtime"
)

func TestByteSliceMD5(t *testing.T) {
	defer xtime.PrintTimeCost("byte slice md5")()

	val := ByteSliceMD5([]byte("ok"))
	assert.Equal(t, "444bcb3a3fcf8389296c49467f27e1d6", val)
}

func TestStrMD5(t *testing.T) {
	defer xtime.PrintTimeCost("string md5")()

	val := StrMD5("ok")
	assert.Equal(t, "444bcb3a3fcf8389296c49467f27e1d6", val)
}

func TestFileMD5(t *testing.T) {
	defer xtime.PrintTimeCost("file md5")()

	val, err := FileMD5("testdata/src.txt")
	assert.Nil(t, err)
	assert.Equal(t, "444bcb3a3fcf8389296c49467f27e1d6", val)
}
