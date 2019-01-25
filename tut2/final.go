// question 9 Complete the program for the output
//Tree maple of height 1.550000 at Price 49.99
// Flower tulip, red at price 1.99
package main

import "fmt"

type Flower struct {
	name  string
	color string
	price float32
}

type Tree struct {
	name   string
	height float64
	price  float32
}

type Plant interface {
	print()
}

func main() {
	gardenStore := [...]Plant{
		Tree{"maple", 1.55, 49.99},
		Flower{"tulip", "red", 1.99}}
	for _,p := range( gardenStore ) {
		p.print()
	}
}

func (t Tree) print() {
	fmt.Printf("Tree %s of height %f at Price %.2f\n", t.name, t.height, t.price)
}

func (f Flower) print() {
	fmt.Printf("Flower %s, %s at price %.2f\n", f.name, f.color, f.price)
}

package main

import "fmt"

type Flower struct {
	name  string
	color string
	price float64
}

type Tree struct {
	name          string
	height        float64
	pricePerMeter float64
}

type Item interface {
	getPrice() float64
}

func main() {
	gardenStore := [2]Item{Tree{"erable", 1.5, 20.00},
		Flower{"tulipe", "red", 2}}
	price := 0.0
	for _, p := range gardenStore {
		price += p.getPrice()
	}

	fmt.Printf("Price: $%6.2f", price)
}

// partial definition of a method (to be completed)
func (t Tree) getPrice() float64 {
	return t.pricePerMeter * t.height
}

// second method (to be completed only if needed)
func (f Flower) getPrice() float64 {
	return f.price

}