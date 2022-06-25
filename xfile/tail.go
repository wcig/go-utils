package xfile

import (
	"io"
	"os"
)

// tail -n lines (默认忽略末尾换行符)
func Tail(filename string, n int) (lines []string, err error) {
	if n <= 0 {
		return nil, nil
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	offsetEnd, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}
	if offsetEnd == 0 {
		return nil, nil
	}
	if offsetEnd > 0 {
		tmp := make([]byte, 1)
		num, err := f.ReadAt(tmp, offsetEnd-1)
		if err != nil {
			return nil, err
		}
		if num > 0 && tmp[0] == '\n' {
			offsetEnd--
		}
	}

	lineNum := 0
	var lineBytes []byte
	for offset := offsetEnd - 1; offset >= 0; offset-- {
		if lineNum >= n {
			break
		}
		tmp := make([]byte, 1)
		num, err := f.ReadAt(tmp, offset)
		if err != nil {
			return nil, err
		}
		if num == 0 {
			break
		}
		if tmp[0] != '\n' {
			lineBytes = append(lineBytes, tmp[0])
		}
		if tmp[0] == '\n' || offset == 0 {
			lineNum++
			reverseBytes(lineBytes)
			lines = append(lines, string(lineBytes))
			lineBytes = []byte{}
		}
	}

	reverseStrings(lines)
	return lines, nil
}

// tail -c bytes (默认忽略末尾换行符)
func TailBytes(filename string, n int) ([]byte, error) {
	if n <= 0 {
		return nil, nil
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	offsetEnd, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}
	if offsetEnd == 0 {
		return nil, nil
	}
	if offsetEnd > 0 {
		tmp := make([]byte, 1)
		num, err := f.ReadAt(tmp, offsetEnd-1)
		if err != nil {
			return nil, err
		}
		if num > 0 && tmp[0] == '\n' {
			offsetEnd--
		}
	}

	size := int64(n)
	if offsetEnd < size {
		size = offsetEnd
	}
	result := make([]byte, size)
	m, err := f.ReadAt(result, offsetEnd-size)
	if err != nil {
		return nil, err
	}
	if int64(m) < size {
		return result[:m], nil
	}
	return result, nil
}

func reverseBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func reverseStrings(data []string) {
	if len(data) == 0 {
		return
	}
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
