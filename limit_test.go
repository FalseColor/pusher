package pusher

import (
	"fmt"
	"testing"
	"time"
)

func Send() int {
	return 0
}

func TestName(t *testing.T) {
	limiter := NewLimiter(1)
	for {
		go limiter.Execute([]byte("888"))
	}
}

// 三次封装
func (l *Limiter) Execute(msg []byte) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		allow := l.Allow()
		if allow {
			defer func() {
				l.Done()
			}()
			time.Sleep(1 * time.Second)
			fmt.Println(string(msg))
			panic("dsad")
		}
	}()

}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("主线程异常关闭")
		}
	}()

	go func() {
		panic("子异常")
	}()
	select {}
}
