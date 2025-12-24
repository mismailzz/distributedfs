package main 

import (
	"github.com/mismailzz/distributedfs/p2p"
)

func main(){

	opts := p2p.TCPTransportOpts {
		ListenAddress: ":3000",
		HandShakeFunc: p2p.NoHandShake,
		DecoderFunc:   &p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(opts)
	tcpTransport.ListenAndAccept()

	select {} // block forever
}