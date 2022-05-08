package main

import (
	"fmt"
	"time"
)
// map在channel是引用类型，会修改到原先的值

var mapChan = make(chan map[string]int, 1)

func send(syncChan chan struct{}){
	countMap := make(map[string]int)
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
			elem["count"]++
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