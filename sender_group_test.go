package pusher

import (
	"fmt"
	"testing"
	"time"
)

func TestSenderGroup_Add(t *testing.T) {
	sender1 := NewSysLogMessageSender("1", "topic1", "tcp", "192.168.1.231:514")
	sender2 := NewSysLogMessageSender("2", "topic2", "tcp", "192.168.1.231:514")
	senders.Add(sender1)
	senders.Add(sender2)
	sender1.Open()
	sender2.Open()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(200 * time.Millisecond)
			senders.Get("1").Send([]byte("message====111===message"))
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(200 * time.Millisecond)
			senders.Get("2").Send([]byte("message====222===message"))
		}
	}()
	for i := 0; i < 10000; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("总体速率：%d，1速率%d，2速率%d \n", senders.GetSpeed(), senders.Get("1").GetSpeed(), senders.Get("2").GetSpeed())

	}
}
func TestSenderGroup_Delete(t *testing.T) {
	sender1 := NewSysLogMessageSender("1", "topic1", "tcp", "192.168.1.231:514")
	sender2 := NewSysLogMessageSender("2", "topic2", "tcp", "192.168.1.231:514")
	senders.Add(sender1)
	senders.Add(sender2)
	sender1.Open()
	sender2.Open()

	time.Sleep(5 * time.Second)
	senders.Delete("1")
	time.Sleep(10 * time.Second)
}
