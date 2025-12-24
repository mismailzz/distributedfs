package p2p

// Some libs needed the hanshake, so lets define here
// HandShakeFunc defines the handshake function signature used by peers.
// It takes a connection and returns an error if the handshake fails.
type HandShakeFunction func(peer *TCPPeer) error

// NoHandShake is a default handshake function that performs no operation.
func NoHandShake(peer *TCPPeer) error {
	return nil
}