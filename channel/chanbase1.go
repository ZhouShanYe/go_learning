package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func receive(syncChan1 chan struct{}, syncChan2 chan struct{}) {
	<-syncChan1
	fmt.Println("Receive a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)

	for {
		if elem, ok := <-strChan; ok{
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}

	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(syncChan1 chan struct{}, syncChan2 chan struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Send:", elem, "[snder]")

		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("send a sync signal. [sender]")
		}
	}

	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second*2)
	close(strChan)
	syncChan2 <- struct{}{}

}

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go receive(syncChan1, syncChan2)
	go send(syncChan1, syncChan2)

	<- syncChan2
	<- syncChan2
}