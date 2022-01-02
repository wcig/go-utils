package xlimiter

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSimpleLimiter(t *testing.T) {
	const num = 5
	limiter := NewSimpleLimiter(2)
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
			limiter.enter()
			defer limiter.leave()
			task(n)
			fmt.Println("over task:", n, start, time.Since(start))
		}(i)
	}
	wg.Wait()

	// Output:
	// add task: 5 2022-01-02 23:08:25.385487 +0800 CST m=+0.000571792
	// add task: 1 2022-01-02 23:08:25.385481 +0800 CST m=+0.000566705
	// add task: 4 2022-01-02 23:08:25.385561 +0800 CST m=+0.000646391
	// add task: 3 2022-01-02 23:08:25.385494 +0800 CST m=+0.000579290
	// add task: 2 2022-01-02 23:08:25.385621 +0800 CST m=+0.000706174
	// over task: 5 2022-01-02 23:08:25.385487 +0800 CST m=+0.000571792 102.641644ms
	// over task: 1 2022-01-02 23:08:25.385481 +0800 CST m=+0.000566705 102.685755ms
	// over task: 3 2022-01-02 23:08:25.385494 +0800 CST m=+0.000579290 206.267575ms
	// over task: 4 2022-01-02 23:08:25.385561 +0800 CST m=+0.000646391 206.261491ms
	// over task: 2 2022-01-02 23:08:25.385621 +0800 CST m=+0.000706174 308.99035ms
}
