package main 

import (
	"os"
	"fmt"
)

type PathKey struct {
	Filename string
	Pathname string 
}

type PathTransformFunction func(string) PathKey

type StoreOpts struct {
	RootDir string
	PathTransformFunc PathTransformFunction 
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store{
	return &Store{
		opts,
	}
}

func (s *Store) writeStream(filename string) error {

	pathKey := s.PathTransformFunc(filename)
	filePath := GetPathWithRoot(s.RootDir, pathKey.Pathname)

	// 1. Create Dir with the required permission 
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}

	// 2. Create file
	filenameWithFullPath := GetFilenameWithFullPath(s.RootDir, pathKey.Pathname, filename)
	f, err := os.Create(filenameWithFullPath)
	if err != nil {
		fmt.Printf("Error while creating file %s:%v", filename, err)
		return err 
	}
	defer f.Close()

	return nil 

}


// Will show a different way to declare the function 
var DefaultPathTransformFunc = func (file string) PathKey {
	
	return PathKey{
		Filename: file, 
		Pathname: file, 
	}
}

func GetPathWithRoot(rootDir string, path string) string {
	return fmt.Sprintf("%s/%s", rootDir, path)
}

func GetFilenameWithFullPath(rootDir string, pathname string, filename string) string {
	return fmt.Sprintf("%s/%s/%s", rootDir, pathname, filename)
}