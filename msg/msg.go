package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

// Processor 使用默认的Protobuf消息处理器
var (
	Processor = protobuf.NewProcessor()
)

func init() {
	Processor.Register(&Heartbeat{})
}
