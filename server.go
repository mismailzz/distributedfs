package main 

import (
	"github.com/mismailzz/distributedfs/p2p"
)

type FileServerOpts struct {
	Transport p2p.Transport // Transport Interface 
	PathTransformFunc PathTransformFunction
	StorageRootDir string 
}

type FileServer struct {
	FileServerOpts
	store *Store
}

func NewFileServer(opts FileServerOpts) *FileServer {
	storeOpts := StoreOpts{
		RootDir: opts.StorageRootDir,
		PathTransformFunc: opts.PathTransformFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store: NewStore(storeOpts),
	}
}

func (f *FileServer) ListenAndAccept() error {
	if err := f.Transport.ListenAndAccept(); err != nil {
		return err
	}
	return nil
}