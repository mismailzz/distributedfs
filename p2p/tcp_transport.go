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
	rpcchan chan RPC // rpc message read chan
}

// TCPTransport constructor
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcchan: make(chan RPC),
	}
}

// ListenAndAccept() implements the Transport Interface
// to start listening and accepting connections
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

	defer conn.Close()

	// As be this step, the connection from remote node has been accepted or established successfully.
	// and passed here to be handled, then we can make it a Peer object. Peer is basically become a second 
	// name of connection in our p2p architecture.
	peer := NewTCPPeer(conn, true) // true indicates this is an outbound connection. As someone came to us, we accepted it and make his connection as outbound.

	log.Printf("New incoming connection from:%+v\n", peer)

	// Perform handshake
	if t.HandShakeFunc != nil { // Check if handshake function is provided
		if err := t.HandShakeFunc(peer); err != nil {
			log.Printf("Handshake failed with %s: %v", conn.RemoteAddr().String(), err) // peer.conn.RemoteAddr() == conn.RemoteAddr() -> will use shortform 
		}
		log.Printf("Handshake successful with %s", conn.RemoteAddr().String()) // peer.conn.RemoteAddr() == conn.RemoteAddr() -> will use shortform
	}

	// Start reading messages from the connection
	rpc := RPC{}
	rpc.Sender = conn.RemoteAddr()
	for { // Loop to read the message from the connection
		err := t.DecoderFunc.Decode(conn, &rpc)
		if err != nil {
			log.Printf("Error decoding message from %s: %v", conn.RemoteAddr().String(), err) // peer.conn.RemoteAddr() == conn.RemoteAddr() -> will use shortform
			return
		
		}
	    // log.Printf("Received message from %s: %s", rpc.Sender.String(), string(rpc.Payload))
		// rather then print - sending RPC messate to channel of that tcp connection
		t.rpcchan <- rpc
	}	

}

// Consume() implements the Transport interface
// to provide the associated channel of a peer, to read the message send by that peer 
func (t *TCPTransport) Consume() <-chan RPC{
	return t.rpcchan
}