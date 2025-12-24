package p2p

import (
	"net"
	"log"
)

const (
	protocolType = "tcp"
)

type TCPTransportOpts struct {
	ListenAddress string
}

type TCPTransport struct {
	TCPTransportOpts
}

// TCPTransport constructor
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	_, err := net.Listen(protocolType, t.ListenAddress) // return listener, error
	if err != nil {
		log.Printf("Error occurred on start listening %s: %v\n", t.ListenAddress, err)
		return err
	}
	log.Printf("listener initialized successfully for address %s", t.ListenAddress)
		
	return nil
}