package p2p 

import (
	"net"
)
 
// TCPPeer represents a remote node in an established TCP connection.
type TCPPeer struct {
	// conn is the underlying TCP connection to the peer (remote node).
	conn net.Conn

	// if outbound is true, then the connection was initiated by this node.
	// It means we dialed out to the remote node and retrieved this connection.
	// if outbound is false, then the connection was accepted from a remote node.
	// It means the remote node dialed us and we accepted the connection.
	outbound bool
}

// NewTCPPeer creates a new TCPPeer instance.
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close() closes the TCP connection to the peer.
func (t *TCPPeer) Close() error {
	return t.conn.Close()
}
