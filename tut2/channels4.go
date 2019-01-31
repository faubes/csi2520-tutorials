package main

import (
	"fmt"
	"unicode"
)

func ProcessLetters(c <-chan rune, out chan<- rune) {
	defer close(out)
	for v := range c {
		out <- unicode.ToUpper(v)
	}
}

func ProcessNumbers(c <-chan int, out chan<- int) {
	sum := 0
	for v := range c {
		sum += v
	}
	out <- sum
	close(out)
}

func main() {
	const str = "CBY B012 has VGA input like it's 1995"
	numbers := make(chan int)
	letters := make(chan rune)

	out_numbers := make(chan int)
	out_letters := make(chan rune)

	go ProcessLetters(letters, out_letters)
	go ProcessNumbers(numbers, out_numbers)

	go func() {
		res := ""
		for char := range out_letters {
			res += string(char)
		}
		fmt.Println(res)
	}()

	go func() {
		val, ok := <-out_numbers
		if !ok {
			panic("Something went very wrong")
		}
		fmt.Println("Sum is ", val)
	}()

	// range syntax:
	// for index, value := range collection
	for _, v := range []rune(str) {
		switch {
		case unicode.IsDigit(v):
			fmt.Println("Digit ", v)
			numbers <- int(v - '0')
		case unicode.IsLetter(v):
			fmt.Println("Letter ", v)
			letters <- v
		default:
			fmt.Println(v)
		}
	}
	close(numbers)
	close(letters)

}
