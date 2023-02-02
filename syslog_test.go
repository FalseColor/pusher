package pusher

import (
	"testing"
)

var senders = SenderGroup{}

func TestSysLog_Send_TCP(t *testing.T) {
	sender, err := NewSysLogMessageSender("sys1", "topic1", "tcp", "192.168.1.231:514")
	if err != nil {
		t.FailNow()
	}
	senders.Add(sender)
	send := senders.Get("sys1")
	for i := 0; i < 10; i++ {
		send.Send([]byte("message====TCP===message"))
	}
}
func TestSysLog_Send_UDP(t *testing.T) {
	sender, err := NewSysLogMessageSender("sys1", "topic1", "udp", "192.168.1.231:514")
	if err != nil {
		t.FailNow()
	}
	senders.Add(sender)
	send := senders.Get("sys1")
	for i := 0; i < 10; i++ {
		send.Send([]byte("message====UDP===message"))
	}
}
