package main

import "fmt"

/*
Compléter le programme suivant afin que la fonction
numberGen genere des entiers entre start et
(start + count),
...
*/
func genNumbers(start, count int, out chan<- int) {
  defer close(out)
  for i:=start; i < start+count; i++ {
    out <- i
  }
}

/*
et la fonction printNumbers imprime les entiers
générés à l’écran:
*/
func printNumbers(in <-chan int, done chan<- bool) {
  for i := range in {
    fmt.Println(i)
  }
  done <- true
}

func main() {

  numbers := make(chan int)
  finished := make(chan bool)
  go genNumbers(16, 7, numbers)
  go printNumbers(numbers, finished)
  <-finished
}
