package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	numFiles = 1000
	fileSize = 1024
	volume   = "/volume"
	dir      = "write-small-files"
)

func main() {
	flag.IntVar(&numFiles, "n", numFiles, "number of files to create")
	flag.IntVar(&fileSize, "s", fileSize, "size of each file")
	flag.StringVar(&volume, "volume", volume, "location of test volume")
	flag.Parse()

	root := filepath.Join(volume, dir)
	if err := os.RemoveAll(root); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(root, 0755); err != nil {
		log.Fatal(err)
	}
	contents := make([]byte, fileSize)
	start := time.Now()
	for i := 0; i < numFiles; i++ {
		file := filepath.Join(root, strconv.Itoa(i))
		if err := ioutil.WriteFile(file, contents, 0644); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("%f\n", time.Since(start).Seconds())
	// Be tidy by clearing up afterwards
	if err := os.RemoveAll(root); err != nil {
		log.Fatal(err)
	}
}
