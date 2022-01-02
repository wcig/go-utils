package xlimiter

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRateLimiter(t *testing.T) {
	const num = 5
	limiter := rate.NewLimiter(2, 1)
	task := func(i int) {
		time.Sleep(100 * time.Millisecond)
	}
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 1; i <= num; i++ {
		go func(n int) {
			start := time.Now()
			fmt.Println("add task:", n, start)
			defer wg.Done()
			_ = limiter.Wait(context.Background())
			task(n)
			fmt.Println("over task:", n, start, time.Since(start))
		}(i)
	}
	wg.Wait()
}
