package xfile

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func init() {
	f, err := os.Create("z.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 1; i <= 100; i++ {
		line := fmt.Sprintf("%04d %s", i, time.Now().Format("2006-01-02 15:04:05.000"))
		_, _ = f.WriteString(line + "\n")
		time.Sleep(time.Millisecond)
	}
}

func BenchmarkTailV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Tail("z.log", 50)
	}
}

func BenchmarkTailV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = TailV2("z.log", 50)
	}
}
