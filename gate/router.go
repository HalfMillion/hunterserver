package gate

import (
	"hunterserver/game"
	"hunterserver/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Heartbeat{}, game.ChanRPC)
}
