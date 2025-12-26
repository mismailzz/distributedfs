package main 

import (
	"testing"
)

func TestWriteStream (t *testing.T){

	file := "example"
	opts := StoreOpts{
		RootDir: "ggnetwork",
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(opts)

	if err := s.writeStream(file); err != nil {
		t.Error(err)
	}
}