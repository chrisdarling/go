package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func md5file(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	h := md5.New()
	io.Copy(h, file)
	return h.Sum(nil)
}

func main() {
	// use walk to walk through a directory and sub directories
	// you get the path and file info
	// return nil if no err occurs
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		b := md5file(path)
		fmt.Printf("%s %x\n", path, b)
		return nil
	})

}
