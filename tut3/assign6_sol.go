package main
import "fmt"
import "os"
import "io"
import "strconv"
import "time"

// part a
func readF2C(fn string, ch chan int) {
	inF, err := os.Open(fn)
	if err != nil {
		fmt.Printf("\nError: could not open file %s", fn)
		return
	}
	var number int
	for {
		_, err := fmt.Fscanf(inF, "%d\n", &number)
		if err != nil {
			if err == io.EOF {
				ch <- -1
			} else {
				fmt.Printf("\nError while reading file %s", fn)
			}
			inF.Close()
			return
		}
		fmt.Printf("%s: <- %d\n", fn, number);
		ch <- number
	}
	inF.Close()
	return
}

// part b
func writeC2F(fn string, ch1 chan int, ch2 chan int) {
	outF, err := os.Create(fn)
	if err != nil {
		fmt.Printf("\nError: could not open file %s", fn)
		return
	}
	var number1 int
	var number2 int
	number1= <-ch1
	number2= <-ch2
	numToWrite := 0
	for {
		if (number1 == -1 && number2 == -1){
			outF.Close()
			return
		} else if number1 == -1 {
			numToWrite = number2
			number2= <-ch2
		} else if number2 == -1 {
			numToWrite = number1
			number1= <-ch1
		} else if number1 >= number2 {
			numToWrite = number2
			number2= <-ch2
		} else if number1 <= number2 {
			numToWrite = number1
			number1= <-ch1
		}
		fmt.Printf("%s: %d =<-\n", fn, numToWrite);
		output := strconv.Itoa(numToWrite)
		_, err := outF.WriteString(output + "\n")
		if err != nil {
			fmt.Printf("\nError while reading file %s", fn)
			outF.Close()
			return
		}
	}
	outF.Close()
	return
}

func main(){
	inputFile1 := "multiples11.txt"
	inputFile2 := "multiples12.txt"
	outputFile := "result.txt"
	
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go readF2C(inputFile1, ch1)
	go readF2C(inputFile2, ch2)
	go writeC2F(outputFile, ch1, ch2)
	time.Sleep(5*time.Second)
}
