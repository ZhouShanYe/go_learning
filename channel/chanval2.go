package main

import (
	"fmt"
	"time"
)
// 太离谱了
type Counter struct{
	count int
}

var mapChan = make(chan map[string]*Counter, 1)

func send(syncChan chan struct{}){
	countMap := map[string]*Counter{
		"count": &Counter{},
	}
	for i := 0; i< 5;i++ {
		mapChan <- countMap
		time.Sleep(time.Millisecond)
		fmt.Printf("The count map: %v. [sender]\n", countMap)
	}

	close(mapChan)
	syncChan <- struct{}{}
}

func receive(syncChan chan struct{}){
	for {
		if elem, ok := <-mapChan; ok{
			counter := elem["count"]
			counter.count++
		} else {
			break
		}
	}

	fmt.Println("Stopped. [receiver]")
	syncChan <- struct{}{}
}

func main() {
	syncChan := make(chan struct{}, 2)
	go send(syncChan)
	go receive(syncChan)

	<- syncChan
	<- syncChan
}