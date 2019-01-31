package main

import (
	"fmt"
)

func main() {

	c := make(chan string)

	c <- "csi2520"
	c <- "csi3104"
	c <- "csi5127"

		msg, ok := <-c
		if !ok {
			panic("ack!")
		} else {
			fmt.Println(msg)
		}
}
