package p2p 

// Peer is an interface that represents the remote node in the p2p network.
type Peer interface {
	Close() error
}

// Every protocol will have the Transport layer for the communication between peer nodes 
// i.e tcp, udp, grpc, etc. 
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC // provides channel to reading the message 
}