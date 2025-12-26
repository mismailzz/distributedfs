package main 

import (
	"os"
	"fmt"
	"io"
	"bytes"
	"log"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

type PathKey struct {
	Filename string
	Pathname string 
}

// Transform func to have a deterministic path and name of file
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

// writeStream - create a file and take the buffer (any kind of data from from io.Reader stream) to write in disk file
func (s *Store) writeStream(filename string, r io.Reader) error {

	pathKey := s.PathTransformFunc(filename)
	filePath := GetPathWithRoot(s.RootDir, pathKey.Pathname)

	// 1. Create Dir with the required permission 
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}

	// 2. Create file
	filenameWithFullPath := GetFilenameWithFullPath(s.RootDir, pathKey.Pathname, pathKey.Filename)
	f, err := os.Create(filenameWithFullPath)
	if err != nil {
		fmt.Printf("Error while creating file %s:%v", filenameWithFullPath, err)
		return err 
	}
	defer f.Close()

	// 3. Write to a file 
	numBytes, err := io.Copy(f, r); 
	if err != nil {
		fmt.Printf("Error while writing to file %s:%v", filenameWithFullPath, err)
		return err 
	}

	fmt.Printf("Written %d bytes to %s\n", numBytes, filenameWithFullPath)

	return nil 
}

// readStread - open a file and take the content to put inside the io.Reader stream 
func (s *Store) readStream(filename string) (io.Reader, error) {

	pathKey := s.PathTransformFunc(filename)
	filenameWithFullPath := GetFilenameWithFullPath(s.RootDir, pathKey.Pathname, pathKey.Filename)

	// 1. Check if file exists or not 
	if !FileExists(filenameWithFullPath) {
		return nil, fmt.Errorf("File %s doesn't exist", filenameWithFullPath)
	}

	// 2. Open file 
	f, err := os.Open(filenameWithFullPath)
	if err != nil {
		return nil, err 
	}

	// 3. Copy file content to buffer
	buf := new(bytes.Buffer)
	numBytes, err := io.Copy(buf, f); 
	if err != nil {
		fmt.Printf("Error while reading file %s:%v", filenameWithFullPath, err)
		return nil, err 
	}	
	
	fmt.Printf("Read %d bytes from %s\n", numBytes, filenameWithFullPath)
	return buf, nil 
}

// delete file after usage
func (s *Store) delete(filename string) error {


	pathKey := s.PathTransformFunc(filename)
	filenameWithFullPath := GetFilenameWithFullPath(s.RootDir, pathKey.Pathname, pathKey.Filename)

	defer func(){
		log.Printf("deleted [%s] from disk", filenameWithFullPath)
	}()

	// 1. Check if file exists or not 
	if !FileExists(filenameWithFullPath) {
		return fmt.Errorf("File %s doesn't exist", filenameWithFullPath)
	}

	// 2. Double deletion attempt for file and dir
	if FileExists(filenameWithFullPath){
		// It can delete only file but not the path  
		if err := os.RemoveAll(filenameWithFullPath); err != nil { return err }
		// due to which we delete the parent directory - which is wierd but workaround
		// need double deletion 
		parentDir := s.RootDir
		if err := os.RemoveAll(parentDir); err != nil { return err }
	}

	return nil 
}

// Will show a different way to declare the function 
var DefaultPathTransformFunc = func (file string) PathKey {
	
	return PathKey{
		Filename: file, 
		Pathname: file, 
	}
}

func CASPathTransformFunc (key string) PathKey {

	// Create determistic hash from same key using SHA1 
	hash := sha1.Sum([]byte(key))
	// Convert the bytes to hex string for hash 
	hashStr := hex.EncodeToString(hash[:])

	// Split the hash string into multiple parts for directory structure (depth levels)
	blocksize := 5
	sliceLen := len(hashStr) / blocksize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
    	from, to := i*blocksize, (i*blocksize)+blocksize
    	paths[i] = hashStr[from:to]
	}	
	
	return PathKey {
		Pathname: strings.Join(paths, "/"), // Join the parts with "/" to form the final path
		Filename: hashStr,
	}
}

func GetPathWithRoot(rootDir string, path string) string {
	return fmt.Sprintf("%s/%s", rootDir, path)
}

func GetFilenameWithFullPath(rootDir string, pathname string, filename string) string {
	return fmt.Sprintf("%s/%s/%s", rootDir, pathname, filename)
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
    	// file doesnt exist 
		return false 
	}
	return true
}