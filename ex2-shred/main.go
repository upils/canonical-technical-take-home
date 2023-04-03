package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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

// shred overwrites the file at the given path with "iterations" times
func shred(path string, iterations uint) error {
	if len(path) == 0 {
		log.Println("No file given, nothing to shred")
		return nil
	}
	if iterations == 0 {
		log.Println("No iterations to run, nothing to do")
		return nil
	}

	fmt.Printf("Will shred %v times %s\n", iterations, path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		log.Println("Given path is a directory, nothing to do")
		return nil
	}

	f, err := os.OpenFile(path, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.CopyN(f, rand.Reader, fileInfo.Size())

	return err
}
