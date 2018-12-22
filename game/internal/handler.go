package internal

import (
	"hunterserver/msg"
	"reflect"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// 向当前模块（注册Heartbet)的消息处理函数
	handler(&msg.Heartbeat{}, handleHeartbeat)
}

func handler(m, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHeartbeat(args []interface{}) {
	// 收到的 Heartbeat 消息
	m := args[0].(*msg.Heartbeat)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息内容
	log.Debug("hell %v", m.Name)

	// 给发送者回应一个 Heartbeat 消息
	a.WriteMsg(&msg.Heartbeat{
		Name: "Gotcha!",
	})
}
