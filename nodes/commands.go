package nodes

import (
	"time"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/serialapi"
)

func (n *Node) On() {
	var send serialapi.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(true)
		cmd.SetNode(n.Id)

		send = cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(100)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Send(send, time.Second)
}

func (n *Node) Off() {
	var send serialapi.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(false)
		cmd.SetNode(n.Id)

		send = cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(0)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Send(send, time.Second)
}

func (n *Node) Level(value float64) {
	var send serialapi.Encodable

	switch {
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(value)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Send(send, time.Second)
}
