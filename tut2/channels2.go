package main

import "fmt"

func ping(c chan<- string, msg string) {
  c <- msg
}

func pong(out <-chan string, in chan<- string) {
  //msg := <-pings
  //pings <- "hello"
  // pongs <- msg
  out <- <-in
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message")
  pong(pings, pongs)
  fmt.Println(<-pongs)
}
