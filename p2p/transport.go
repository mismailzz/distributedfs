package p2p 


// Every protocol will have the Transport layer for the communication between peer nodes 
// i.e tcp, udp, grpc, etc. 
type Transport interface {
	ListenAndAccept() error
}