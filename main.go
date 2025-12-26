package main 

import (
	"fmt"
	"github.com/mismailzz/distributedfs/p2p"
)

func main(){

	// 1. Initialize the Transport to TCPTransport 
	opts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandShakeFunc: p2p.NoHandShake,
		DecoderFunc:   &p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(opts)

	// 2.  Initialize the FileServer with Store and Transport
	serverOpts := FileServerOpts{
		Transport: tcpTransport, // interesting because in NewFileServer we are'nt initializing there
		PathTransformFunc: CASPathTransformFunc,
		StorageRootDir: "3000_network",
	}

	// 3. Create the FileServer and start listening
	server := NewFileServer(serverOpts)

	if err := server.ListenAndAccept(); err != nil {
		fmt.Errorf("Error happend starting Server: %v", err)
	}
	
	select {} // block forever
}

// This can be removed by writing the test cases 
func mainTCPTransport() {

	opts := p2p.TCPTransportOpts {
		ListenAddress: ":3000",
		HandShakeFunc: p2p.NoHandShake,
		DecoderFunc:   &p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(opts)
	tcpTransport.ListenAndAccept()

	// Go routine: to read the message from the channel being sent 
	go func() {
		for {
			msg := <- tcpTransport.Consume()
			fmt.Printf("Message Recieved: %+v\n", msg)
		}
	}()

	select {} // block forever
}