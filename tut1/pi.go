package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"
)

type Point struct {
  x, y float64
}

func generatePoint() Point {
  return Point{rand.Float64(), rand.Float64()}
}

func (p Point) distanceFromOrigin() float64 {
  return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p Point) isInQuadrant() bool {
  return p.distanceFromOrigin() < 1
}

func monteCarloPi(print_iterations int) func() float64 {
  in := 0
  out := 0
  return func() float64 {
    p := generatePoint()
    //fmt.Println(p)
    switch p.isInQuadrant() {
    case true:
        //fmt.Println("In")
        in++
    case false:
        //fmt.Println("Out")
        out++
    }
    if (in + out) % print_iterations == 0 {
      fmt.Printf("%d in, %d out\n", in, out)
    }
    return 4*float64(in)/float64(in+out)
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())
  const iterations int = 20000
  m := monteCarloPi(iterations / 10)
  for i := 0; i < iterations; i++ {
    m()
  }
  result := m()
  fmt.Printf("Approx: %d, Pi: %d\n", result, math.Pi)
  fmt.Printf("Difference: %d\n", math.Abs(result - math.Pi))
}
