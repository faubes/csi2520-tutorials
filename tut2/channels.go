package main

import (
	"fmt"
)

func main() {

	c := make(chan string)
	c <- "csi2520"
	msg, ok := <-c
	if !ok {
		panic("ack!")
	} else {
		fmt.Println(msg)
	}
}
