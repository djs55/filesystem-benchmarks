package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

const (
	image = "djs55/write-small-files:latest"
)

func main() {
	volume, err := filepath.Abs("./volume")
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&volume, "volume", volume, "path for shared files")
	flag.Parse()

	if err := os.MkdirAll(volume, 0755); err != nil {
		log.Fatal(err)
	}
	if err := docker("pull", image); err != nil {
		log.Fatal(err)
	}
	hostWrites := path.Join(volume, "host-writes")
	if err := os.MkdirAll(hostWrites, 0755); err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(hostWrites)
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		i := 0
		defer wg.Done()
		defer func() {
			log.Printf("Wrote %d small files on the host in the background\n", i)
		}()
		for {
			name := path.Join(hostWrites, strconv.Itoa(i))
			if err := ioutil.WriteFile(name, nil, 0644); err != nil {
				log.Fatal(err)
			}
			i++
			select {
			case <-stop:
				return
			default:
			}
		}
	}()
	if err := docker("run", "-v", volume+":/volume", image); err != nil {
		log.Fatal(err)
	}
	close(stop)
}

func docker(args ...string) error {
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("docker %s\n", strings.Join(args, " "))
	return cmd.Run()
}
