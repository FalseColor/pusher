package pusher

type MessageSender interface {
	GetName() string
	Connect() error
	//Send(msg []byte) error

	SendAsync(message []byte)
	Open() error // 状态计算
	Close()      // 停止计算状态
	GetSpeed() uint64
	//Test()
}
