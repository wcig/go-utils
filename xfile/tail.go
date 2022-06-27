package xfile

import (
	"bytes"
	"io"
	"os"
)

const (
	minTailBufferSize     = 16
	defaultTailBufferSize = 4096
)

// tail -n lines (包括行末尾换行符)
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

		isFirst := offset == offsetEnd-1
		if tmp[0] != '\n' || isFirst {
			lineBytes = append(lineBytes, tmp[0])
		}
		if tmp[0] == '\n' || offset == 0 {
			if isFirst {
				continue
			}
			lineNum++
			reverseBytes(lineBytes)
			lines = append(lines, string(lineBytes))
			if tmp[0] == '\n' {
				lineBytes = []byte{'\n'}
			} else {
				lineBytes = []byte{}
			}
		}
	}
	if len(lineBytes) > 0 && lineNum < n {
		lines = append(lines, string(lineBytes))
	}

	reverseStrings(lines)
	return lines, nil
}

func TailV2(filename string, n int) (lines []string, err error) {
	return TailV2WithBuf(filename, n, defaultTailBufferSize)
}

func TailV2WithBuf(filename string, n int, bufSize int) (lines []string, err error) {
	if n <= 0 {
		return nil, nil
	}

	if bufSize < minTailBufferSize {
		bufSize = minTailBufferSize
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	isLastByteLineBreak := false
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
			isLastByteLineBreak = true
			offsetEnd--
		}
	}

	var (
		buf       []byte
		lineBytes []byte
		lineNum   = 0
	)

	for offset := offsetEnd; offset > 0 && lineNum < n; {
		if offset >= int64(bufSize) {
			buf = make([]byte, bufSize)
			offset -= int64(bufSize)
		} else {
			buf = make([]byte, offset)
			offset = 0
		}

		_, err = f.ReadAt(buf, offset)
		if err != nil {
			return nil, err
		}

		for {
			index := bytes.LastIndexByte(buf, '\n')
			if index >= 0 {
				tmp := make([]byte, len(buf)-index)
				copy(tmp, buf[index:])
				lineBytes = append(tmp, lineBytes...)
				buf = buf[:index]

				lines = append(lines, string(lineBytes[1:]))
				lineBytes = []byte{}
				lineNum++
				if lineNum >= n {
					break
				}
			} else {
				lineBytes = append(buf, lineBytes...)
				break
			}
		}

		if offset <= 0 && len(lineBytes) > 0 {
			lines = append(lines, string(lineBytes))
			lineBytes = []byte{}
		}
	}

	reverseStrings(lines)
	for i := range lines {
		if i != len(lines)-1 || isLastByteLineBreak {
			lines[i] += "\n"
		}
	}
	return lines, nil
}

// tail -c bytes (包括行末尾换行符)
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
