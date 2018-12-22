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
	log.Debug("收到消息")

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

/*
   序号 字节  字节数 类型     意义     描述
    0  0000    4   Int     size    包大小 后接12个字节
    4  00      2   Short   TP
    6  0       1   Byte    1
    7  00      2   Short   gameID
    9  0000    4   Int     cmd
    13 00      2   Short   subCmd
    15 0       1   Byte    0
*/
