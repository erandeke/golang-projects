package main

import "fmt"

func main() {

	mychannel := make(chan int)

	go sendValuesOverChannel(mychannel)

	for i := 0; i < 6; i++ {
		fmt.Println(<-mychannel)
	}

}

func sendValuesOverChannel(mychannel chan int) {
	for i := 0; i < 6; i++ {
		mychannel <- i //sending value
	}
	close(mychannel)

}
