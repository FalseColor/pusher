package pusher

import (
	"fmt"
	"time"
)

type SysLog struct {
	name        string
	count       uint64
	speed       uint64
	topic       string
	network     string
	address     string
	stopChannel chan int
	status      int // 0关闭，1开启
	workerCount int
	limiter     *Limiter
}

func (s *SysLog) Send(msg []byte) error {
	return nil
}
func (s *SysLog) SendAsync(msg []byte) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		allow := s.limiter.Allow()
		if allow {
			defer func() {
				s.limiter.Done()
			}()
			s.Send(msg)
		}
	}()
}
func (s *SysLog) Connect() error {

	return nil
}
func (s *SysLog) Open() error {
	err := s.Connect()
	if err != nil {
		return err
	}
	s.stopChannel = make(chan int)
	s.limiter = NewLimiter(s.workerCount)
	s.status = 1
	go func() {
		for {
			select {
			case <-s.stopChannel:
				return
			// 退出
			default:
				// 5 秒统计一次
				s.speed = s.count / 5
				s.count = 0
			}
			fmt.Println(s.name)
			// 并发可能会出现漏统计，不计
			time.Sleep(5 * time.Second)
		}
	}()
	return nil

}
func (s *SysLog) Close() {
	s.stopChannel <- 1

}

func (s *SysLog) GetSpeed() uint64 {
	return s.speed
}
func (s *SysLog) GetName() string {
	return s.name
}

func NewSysLogMessageSender(name string, topic string, network network, address string, workerCount int) MessageSender {
	return nil
}
