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
	HandShakeFunc HandShakeFunction
	DecoderFunc   Decoder
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
	go t.startAcceptLoopForConnection(t.listener)

	return nil
}

// Accept incoming connections in a loop
func (t *TCPTransport) startAcceptLoopForConnection(listener net.Listener) {
	for {
		conn, err := listener.Accept() // blocking system call, will wait/halt here until a new connection comes in
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go t.handleNewConnection(conn)
	}
}

// Handle new incoming connection - every connection gets its own goroutine to handle its connection
func (t *TCPTransport) handleNewConnection(conn net.Conn) {
	log.Printf("New connection established from %s", conn.RemoteAddr().String())

	// Perform handshake
	if t.HandShakeFunc != nil { // Check if handshake function is provided
		if err := t.HandShakeFunc(conn); err != nil {
			log.Printf("Handshake failed with %s: %v", conn.RemoteAddr().String(), err)
		}
		log.Printf("Handshake successful with %s", conn.RemoteAddr().String())
	}

	rpc := RPC{}
	rpc.Sender = conn.RemoteAddr()
	for { // Loop to read the message from the connection
		err := t.DecoderFunc.Decode(conn, &rpc)
		if err != nil {
			log.Printf("Error decoding message from %s: %v", conn.RemoteAddr().String(), err)
			continue
		}
		log.Printf("Received message from %s: %s", rpc.Sender.String(), string(rpc.Payload))
	}

}


// TODO: Close connectin and listener methods