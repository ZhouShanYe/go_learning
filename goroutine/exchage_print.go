package main

import (
	"fmt"
	// "time"
)

func print1_1(chan0 chan int, chan1 <-chan int, chan2 chan<- int){
	for i:=0; i< 10;i++{
		<-chan1
		fmt.Print(1)
		chan2<- 2
	}
}

func print2_1(chan0 chan<- int, chan1 chan<- int, chan2 <-chan int){
	for i:=0; i< 10;i++{
		<-chan2
		fmt.Print(2)
		chan1 <- 1
	}
	chan0 <- 0
}

func printWithBufferChannel(){
	chan0 := make(chan int, 1)
	chan1 := make(chan int, 1)
	chan2 := make(chan int)
	chan1 <- 1
	go print1_1(chan0, chan1, chan2)
	go print2_1(chan0, chan1, chan2)
	<-chan0
}

func print1_2(chan_print1 chan int, chan_print2 chan int){
	for i := 0;i < 10; i++{
		chan_print1<-0
		fmt.Print(1)
		<-chan_print2
	}
}

func print2_2(chan_print1 chan int, chan_print2 chan int, exit chan <- int){
	for i := 0;i < 10; i++{
		<-chan_print1
		fmt.Print(2)
		chan_print2<-0
	}
	exit<-0
}

func printWithoutBufferChannel(){
	exit := make(chan int)
	chan_print1 := make(chan int)
	chan_print2 := make(chan int)

	go print1_2(chan_print1, chan_print2)
	go print2_2(chan_print1, chan_print2, exit)
	<-exit
}

func printNumber(chanInt chan int){
	
}
func printNumberWithNGoroutine(){
	n := 3

	chan_int := make(chan int)
	for i := 0; i < n; i++{
		go printNumber()
	}
}
func main(){
	// printWithBufferChannel()
	// printWithoutBufferChannel()
	printNumberWithNGoroutine()
}