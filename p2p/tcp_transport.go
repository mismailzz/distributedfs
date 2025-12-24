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
	listener net.Listener
}

// TCPTransport constructor
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

// Start listening and accepting connections
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen(protocolType, t.ListenAddress)
	if err != nil {
		log.Printf("Error occurred on start listening %s: %v\n", t.ListenAddress, err)
		return err
	}
	log.Printf("Listening on address %s", t.ListenAddress)
	go t.AcceptLoopForConnections(t.listener)

	return nil
}

// Accept incoming connections in a loop
func (t *TCPTransport) AcceptLoopForConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go t.HandleConnection(conn)
	}
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	log.Printf("New connection established from %s", conn.RemoteAddr().String())
}


// TODO: Close connectin and listener methods