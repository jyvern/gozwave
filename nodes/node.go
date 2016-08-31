package nodes

import (
	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/functions"
	"github.com/stampzilla/gozwave/serialapi"
)

func New(address byte, connection *serialapi.Connection) *node {
	return &node{
		id:         address,
		connection: connection,
	}
}

type node struct {
	id byte

	ProtocolInfo        *functions.FuncGetNodeProtocolInfo
	ManufacurerSpecific *commands.CmdManufacturerSpecific

	connection *serialapi.Connection
	awake      chan struct{}
	identified bool
}

func (n *node) Id() byte {
	return n.id
}

func (n *node) isAwake() chan struct{} {
	c := make(chan struct{})
	close(c)
	return c
}

func (n *node) Identify() {
	logrus.Warnf("Started identification on node %d", n.Id())

	for {
		<-n.isAwake()

		// All manufacturer specific information such as vendor, vendors product ID and product type
		if n.ManufacurerSpecific == nil {
			err := n.RequestManufacturerSpecific()
			if err != nil {
				continue
			}
		}

		//<-self.Connection.SendRaw([]byte{functions.GetNodeProtocolInfo, byte(index + 1)}) // Request node information
		//		nodeinfo := self.WaitForGetNodeProtocolInfo()

		//<-self.Connection.SendRaw([]byte{functions.IsFailedNode, byte(index + 1)}) // Request is failed node
		//	<-self.SendRaw([]byte{0xa0, byte(index + 1)}) // Request ?

		// All firmware and Z-Wave version information
		// All switching and reporting capabilities including current switching states and sensor values.
		// Number and maximum size of association groups
		// Configuration values if known
		// The Gateways will then do some initial setup:
		// If needed a certain predefined wakeup interval is set for battery operated devices
		// If available and requested certain standard defaults behaviours are configured in the device
		// If association groups are available the gateway will always put its own Node ID into these association groups in order to be informed about status changes and to be prepared for using associations for scene switching
		// The user can change all values. However it needs to be clear that particularly removing the gateways Node ID from the association groups may cause malfunctions of the gateway.

		n.identified = true
		return
	}
}
