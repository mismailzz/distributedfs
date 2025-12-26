package main 

import (
	"testing"
	"bytes"
	"fmt"
	"io/ioutil"

)

func TestWriteStream (t *testing.T){

	file := "example"
	data := []byte("hello world!")

	opts := StoreOpts{
		RootDir: "ggnetwork",
		PathTransformFunc: CASPathTransformFunc, //DefaultPathTransformFunc,
	}

	s := NewStore(opts)

	// Write 
	if err := s.writeStream(file, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	// Read
	r, err := s.readStream(file)
	if err != nil { 
		t.Errorf("readStream %s\n", err)
	}

	b, _ := ioutil.ReadAll(r)
	fmt.Println(string(b))

	if string(b) != string(data){
		t.Errorf("want %s have %s", data, b)
	}

	// Delete
	if err = s.delete(file); err != nil {
		t.Error(err)
	} 

	// Check if Deletion happened 
	pathKey := s.PathTransformFunc(file)
	filenameWithFullPath := GetFilenameWithFullPath(s.RootDir, pathKey.Pathname, pathKey.Filename)
	if FileExists(filenameWithFullPath){
		t.Error(err)
	}

}