package p2p 

import (
	"io"
)

// Any Protocol can define its own Decoder to decode bytes into Message.
type Decoder interface {
	Decode(conn io.Reader, rpc *RPC) error
}

// DefaultDecoder is a basic implementation of the Decoder interface.
type DefaultDecoder struct {}
func (d *DefaultDecoder) Decode(conn io.Reader, rpc *RPC) error {

	buf := make([]byte, 1024) // buffer to hold incoming data
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	
	rpc.Payload = buf[:n]
	
	return nil
}