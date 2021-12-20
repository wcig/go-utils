package xmd5

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// ByteSliceMD5 cal byte slice md5
func ByteSliceMD5(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

// StrMD5 cal string md5
func StrMD5(data string) string {
	return ByteSliceMD5([]byte(data))
}

// StrMD5 cal file md5
func FileMD5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	r := bufio.NewReader(f)
	h := md5.New()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
