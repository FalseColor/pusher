package pusher

type MessageSender interface {
	GetName() string
	Connect() error
	Send(msg []byte) error
	Open() error // 状态计算
	Close()      // 停止计算状态
	GetSpeed() uint64
}
