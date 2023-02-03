package pusher

import (
	"fmt"
	"log/syslog"
	"time"
)

type SysLog struct {
	name        string
	count       uint64
	speed       uint64
	topic       string
	network     string
	address     string
	dial        *syslog.Writer
	stopChannel chan int
	status      int // 0关闭，1开启
}

func (s *SysLog) Send(msg []byte) error {
	if s.status == 0 {
		return nil
	}
	if s.dial == nil {
		s.Connect()
	}
	n, err := s.dial.Write(msg)
	s.count += uint64(n)
	return err
}
func (s *SysLog) Connect() error {
	dial, err := syslog.Dial(s.network, s.address, syslog.LOG_INFO, s.topic)
	if err != nil {
		return err
	}
	s.dial = dial
	return nil
}
func (s *SysLog) Open() error {
	err := s.Connect()
	if err != nil {
		return err
	}
	s.stopChannel = make(chan int)
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
	s.dial.Close()
}

func (s *SysLog) GetSpeed() uint64 {
	return s.speed
}
func (s *SysLog) GetName() string {
	return s.name
}

func NewSysLogMessageSender(name string, topic string, network network, address string) (MessageSender, error) {
	log := SysLog{
		name:    name,
		topic:   topic,
		network: string(network),
		address: address,
	}
	err := log.Open()
	return &log, err
}
