package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"
)

const (
	volume = "/volume"
	app    = "my-app"
)

func main() {
	fmt.Println("# npx create-react-app => npm start benchmark")
	if _, err := os.Stat(path.Join(volume, app)); os.IsNotExist(err) {
		bootstrap()
	}
	fmt.Println("# app has been bootstrapped")
	fmt.Println("# iteration, time/seconds")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d, %f\n", i, measure().Seconds())
	}
	fmt.Println("# done")
}

func measure() time.Duration {
	start := time.Now()
	cmd := exec.Command("npm", "start")
	cmd.Dir = path.Join(volume, app)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("# %s\n", line)
		}
	}()
	// https://github.com/facebook/create-react-app/issues/8688
	cmd.Env = append(os.Environ(), "CI=true")
	// Create all subprocesses in a group so we can kill them later
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	defer cmd.Process.Wait()
	defer syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("# %s\n", line)
		if strings.Contains(line, "Compiled successfully!") {
			return time.Since(start)
		}
	}
	log.Fatal("never saw 'Compiled successfully!")
	return time.Duration(0)
}

func bootstrap() {
	cmd := exec.Command("npx", "create-react-app", app)
	cmd.Dir = volume
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
