package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrEmptyPath    error = fmt.Errorf("no file given, nothing to shred")
	ErrNoIterration error = fmt.Errorf("no iterations to run, nothing to do")
	ErrPathIsDir    error = fmt.Errorf("given path is a directory, nothing to do")
)

func main() {
	var iterationsFlag uint
	flag.UintVar(&iterationsFlag, "n", 0, "number of iterations")
	var pathFlag string
	flag.StringVar(&pathFlag, "path", "", "path of the file to shred")
	flag.Parse()

	err := shred(pathFlag, iterationsFlag)
	if err != nil {
		log.Fatal(err)
	}
}

// shred overwrites the file at the given path, "iterations" times
func shred(path string, iterations uint) error {
	if len(path) == 0 {
		return ErrEmptyPath
	}
	if iterations == 0 {
		return ErrNoIterration
	}

	fmt.Printf("Will shred %v times %s\n", iterations, path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return ErrPathIsDir
	}

	f, err := os.OpenFile(path, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	var i uint
	for i = 0; i < iterations; i++ {
		_, err = f.Seek(0, 0)
		if err != nil {
			return err
		}
		_, err = io.CopyN(f, rand.Reader, fileInfo.Size())
		if err != nil {
			return err
		}
	}

	return nil
}
