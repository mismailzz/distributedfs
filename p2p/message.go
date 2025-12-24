package p2p 

import (
	"net"
)

// RPC represents a remote procedure call message, which is basically just a message
// being transmitted between the two nodes/peers in the P2P network.
type RPC struct {
	Sender net.Addr
	Payload []byte
}