package xfile

import (
	"bufio"
	"os"
)

// head -n lines
func Head(filename string, n int) (lines []string, err error) {
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

	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i++
		if i > n {
			break
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
