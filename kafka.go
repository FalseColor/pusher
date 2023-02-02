package pusher

import (
	"github.com/Shopify/sarama"
	"time"
)

type Kafka struct {
	name        string
	count       uint64
	speed       uint64
	topic       string
	address     []string
	username    string
	password    string
	dial        sarama.SyncProducer
	stopChannel chan int
	status      int // 0关闭，1开启
}

func (k *Kafka) Send(message []byte) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = k.topic
	msg.Value = sarama.ByteEncoder(message)
	// 发送消息
	_, _, err := k.dial.SendMessage(msg)
	return err
}
func (k *Kafka) Connect() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = false                  // 成功交付的消息不会在success channel返回
	config.Net.SASL.Enable = true
	config.Net.SASL.User = k.username
	config.Net.SASL.Password = k.password
	// 连接kafka
	client, err := sarama.NewSyncProducer(k.address, config)
	if err != nil {
		return err
	}
	k.dial = client
	return nil
}
func (k *Kafka) Open() {
	k.status = 1
	k.stopChannel = make(chan int)
	go func() {
		for {
			select {
			case <-k.stopChannel:
				return
			// 退出
			default:
				// 5 秒统计一次
				k.speed = k.count / 5
				k.count = 0
			}
			// 并发可能会出现漏统计，不计
			time.Sleep(5 * time.Second)
		}
	}()

}
func (k *Kafka) Close() {
	k.stopChannel <- 1
	k.dial.Close()
}
func (k *Kafka) GetSpeed() uint64 {
	return k.speed
}
func (k *Kafka) GetName() string {
	return k.name
}
func NewKafkaMessageSender(name string, address []string, username string, password string, topic string) (MessageSender, error) {
	kafka := Kafka{
		name:     name,
		topic:    topic,
		address:  address,
		username: username,
		password: password,
	}
	err := kafka.Connect()
	kafka.Open()
	return &kafka, err
}
