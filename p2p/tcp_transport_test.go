package p2p

import (
	"testing"
)

func TestListenAndAccept(t *testing.T) {
	opts := TCPTransportOpts {
		ListenAddress: ":3000", // Same as 0.0.0.0:3000 as anyone (local + network + internet*) can connect to this address
	}
	tcpTransport := NewTCPTransport(opts)
	if err := tcpTransport.ListenAndAccept(); err != nil {
		t.Error(err)
	}
}