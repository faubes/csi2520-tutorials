package main

import (
	"fmt"
)

// empty interface
// all types implement I since it has no methods
type I interface{}

// print function uses switch and i.(type)
func print(i I) {
	switch v := i.(type) {
	case []int:
		// v is an []int
		fmt.Printf("Type: %T, len: %d, cap: %d, Value: %v\n", v, len(v), cap(v), v)
	case string:
		// v is a string
		fmt.Printf("Type: %T, Value: %s\n", v)
	default:
		fmt.Printf("Type: %T, Value: %v\n", v, v)
	}
}

// splits slice into first half, second half
func splitSlice(slice []int) (a []int, b []int) {
	mid := len(slice) / 2
	return slice[:mid], slice[mid:]
}

// returns pointers to min/max elements in a slice
func minMax(slice []int) (min *int, max *int) {
  min, max = &slice[0], &slice[0]
	// range function returns index i, value v
  for i, v := range slice {
      if v < *min {
        min = &slice[i]
      }
      if v > *max {
        max = &slice[i]
      }
    }
  return
}

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}
  print(array1)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// create slice on array
	slice2 := array1[:]
	print(slice1)
  print(slice1[len(slice1)-1:])
	first, second := splitSlice(slice1)
	print(first)
	print(second)
	//careful: first and second slice
	// point to the same underlying array as slice1:
	first[0] = 42
	// changing first[0] also changes slice1[0]
	print(slice1)
	// double pointer return
  min, max := minMax(slice1)
	// assigning through a pointer
  *min = -100

  var pair []int
	// using append to grow a slice
	// underlying array capacity increased if needed
  pair = append(pair, *min)
  pair = append(pair, *max)
  print(pair)

}
