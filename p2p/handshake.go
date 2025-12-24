package p2p

import (
	"net"
)

// Some libs needed the hanshake, so lets define here
// HandShakeFunc defines the handshake function signature used by peers.
// It takes a connection and returns an error if the handshake fails.
type HandShakeFunction func(conn net.Conn) error

// NoHandShake is a default handshake function that performs no operation.
func NoHandShake(conn net.Conn) error {
	return nil
}