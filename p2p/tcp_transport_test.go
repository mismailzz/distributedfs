package p2p

import (
	"testing"
)

func TestListenAndAccept(t *testing.T) {
	opts := TCPTransportOpts {
		ListenAddress: ":3000",
	}
	tcpTransport := NewTCPTransport(opts)
	if err := tcpTransport.ListenAndAccept(); err != nil {
		t.Error(err)
	}
}