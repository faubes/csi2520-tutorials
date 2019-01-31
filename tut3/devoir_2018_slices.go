package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"bufio"
	"os"
)

func AbsDiff(sliceA, sliceB []float32) (res []float32, err error) {
	fmt.Printf("Slice A: %v\n", sliceA)
	fmt.Printf("Slice B: %v\n", sliceB)

	if len(sliceA) == 0 {
		sliceA = make([]float32, len(sliceB))
		copy(sliceA, sliceB)
	}

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

func getSlice() ([]float32, error) {
	var s []float32
	fmt.Println("Enter another slice of floats:")
	scanner := bufio.NewScanner(os.Stdin) // create scanner on stdin
  scanner.Scan() // read single line (default split) from scanner
  input := scanner.Text() // store string in scanner
//  fmt.Println(input)
	tokens := strings.Split(input, " ") // split string
	for _, v := range tokens {
		f, ok := strconv.ParseFloat(v, 32) // parse to float
		if ok != nil {
			return []float32{}, ok
		}
		s = append(s, float32(f)) // add float to slice
	}
	return s, nil
}

func quit() bool {
	var c rune
	fmt.Print("q to quit (anything else to continue): ")
	fmt.Scanf("%c", &c)
	if c == 'q' {
		return true
	} else {
		return false
	}
}

func main() {
	var slice1, slice2 []float32
	var ok error
	for {
		fmt.Printf("Previous slice: %v\n", slice1)
		for {
			slice2, ok = getSlice()
			if ok == nil {
				break
			} else {
				fmt.Println(ok)
			}
		}
		fmt.Printf("New slice: %v\n", slice2)
		result, err := AbsDiff(slice1, slice2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result: %v\n", result)
			slice1 = make([]float32, len(slice2))
			copy(slice1, slice2)
		}
		if quit() {
			break
		}
	}
}
