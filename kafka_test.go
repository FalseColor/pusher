package pusher

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

func TestNewKafkaMessageSender(t *testing.T) {

	//sender := NewKafkaMessageSender("http", []string{"192.168.1.231:9092"}, "client1", "pass1", "top111", 1)
	//err := sender.Connect()
	//sender.SendAsync([]byte("888"))
	//senders.Add(sender)
	client, err := sarama.NewClient([]string{"192.168.1.231:9092"}, NewTestConfig())
	if err != nil {
		t.FailNow()
	}
	topics, err := client.Topics()
	fmt.Println(topics)
}

func NewTestConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息不会在success channel返回
	config.Net.SASL.Enable = true
	config.Net.SASL.User = "client1"
	config.Net.SASL.Password = "pass1"
	config.Version = sarama.V3_2_3_0
	return config
}
