package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

// Processor 使用默认的Protobuf消息处理器
var Processor protobuf.Processor

func init() {
	// json.NewProcessor
}

// Heartbeat 心跳处理
type Heartbeat struct {
	Name string
}
