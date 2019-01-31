package main

import (
	"fmt"
  "math"
	"math/rand"
  "errors"
)

func RandomArray(len int) []float32 {
	array := make([]float32, len)
	for i := range array {
		array[i] = rand.Float32()
	}
	return array
}

func AbsDiff(sliceA, sliceB []float32) (res []float32, err error) {
//  fmt.Printf("Slice A: %v\n", sliceA)
//  fmt.Printf("Slice B: %v\n", sliceB)
  if len(sliceA) != len(sliceB) {
		return nil, errors.New("Slice lengths do not match")
	}
	var dist []float32
	for i := range sliceA {
		difference := float64(sliceA[i] - sliceB[i])
		result := float32(math.Abs(difference))
		dist = append(dist, result)
	}
	return dist, nil
}

func Process(a []float32, out chan<- float32) {
  mid := len(a) / 2
  dist, err := AbsDiff(a[:mid], a[mid:])
  if err != nil {
    panic("Whoa!")
  }
  var sum float32 = 0.0
  for _, v := range dist {
    sum += v
  }
  out <- sum
}

func main() {
	rand.Seed(100) // use this seed value
	out := make(chan float32)
	defer close(out)
	for i := 0; i < 1000; i++ {
		a := RandomArray(2 * (50 + rand.Intn(50)))
		go Process(a, out)
	}
  var sum float32 = 0.0
  for i := 0; i < 1000; i++ {
    sum += <-out
  }

  fmt.Println(sum)
}
