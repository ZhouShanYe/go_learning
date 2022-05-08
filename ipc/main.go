package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
)

func testPipe() {
	cmd := exec.Command("echo", "-n", "My first command comes from golang.s")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// output := make([]byte, 30)
	// n, err := stdout.Read(output)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", output[:n])

	var outputBuf bytes.Buffer
	for {
		tempOutput := make([]byte, 5)
		n, err := stdout.Read(tempOutput)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		if n > 0 {
			outputBuf.Write(tempOutput[:n])
		}
	}
	fmt.Printf("%s\n", outputBuf.String())
}

func testSignal() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s...[sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("Set notification for %s...[sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv1\n")
		wg.Done()
	}()
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv2\n")
		wg.Done()
	}()

	wg.Wait()
	signal.Stop(sigRecv1)
	close(sigRecv1)
}

func main() {
	fmt.Println("ipc.go")

	testPipe()
	testSignal()
}
