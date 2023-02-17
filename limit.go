package pusher

type Limiter struct {
	rate chan int
}

func NewLimiter(num int) *Limiter {
	limiter := Limiter{rate: make(chan int, num)}
	return &limiter
}

func (l *Limiter) Allow() bool {
	select {
	case l.rate <- 1:
		return true
	default:
		return false
	}
}
func (l *Limiter) Done() {
	<-l.rate
}
