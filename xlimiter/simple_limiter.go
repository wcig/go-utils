package xlimiter

// concurrent limiter
type SimpleLimiter struct {
	sem chan struct{}
}

func NewSimpleLimiter(num int64) *SimpleLimiter {
	return &SimpleLimiter{sem: make(chan struct{}, num)}
}

func (sl *SimpleLimiter) enter() {
	sl.sem <- struct{}{}
}

func (sl *SimpleLimiter) leave() {
	<-sl.sem
}
