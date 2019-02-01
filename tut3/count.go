/* fun with channels */

package main

import (
	"fmt"
	"time"
)


func count(name string, in <-chan int, out chan<- int,
		stop int, done chan bool) {
	for {
  i, ok := <-in
	if !ok || i > stop {
		done <- true
		close(out)
		return
	}
	fmt.Println(name, i)
	time.Sleep(time.Millisecond*500)
	i++
	out <- i
	}
}

func main() {
	joe := make(chan int)
	jane := make(chan int)
	juliet := make(chan int)
	done := make(chan bool)
	go func() {
		joe <- 1
	}()
	go count("Joe", joe, jane, 15, done)
	go count("Jane", jane, juliet, 15, done)
	go count("Juliet", juliet, joe, 15, done)
	<-done
	<-done
	<-done
}
