package main

import "fmt"

func produce(dataChannel chan int) {
	for i := 0; i < 10; i++ {
		dataChannel <- i
	}
}

func something() {
	dataChannel := make(chan int)

	go produce(dataChannel)
	go produce(dataChannel)
	go produce(dataChannel)

	for i := 0; i < 30; i++ {
		data := <-dataChannel
		fmt.Printf("%v", data)
	}
	fmt.Print("\n")
}

func main() {
	something()
	something()
	something()
	something()
}
