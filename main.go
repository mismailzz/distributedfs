package main 

import (
	"fmt"
	"github.com/mismailzz/distributedfs/p2p"
)

func main(){

	fmt.Println("HelloWorld")
	
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