package main

import (
	"fmt"
	"math"
)

// define interface shape: has two methods
type shape interface {
	name() string
	area() float64
	perimeter() float64
}

// declare types rectangle, circle
type rect struct {
	s string
	height, width float64
}

type tri struct {
	s string
	height, base float64
}

type circ struct {
	s string
	radius float64
}

// name getters
func (r rect) name() string { return r.s }
func (t tri) name() string { return t.s }
func (c circ) name() string { return c.s }

// method invoked by ex: r1.area()
func (r rect) area() float64 {
	return r.height * r.width
}

// r1.perimeter()
func (r rect) perimeter() float64 {
	return 2*r.height + 2*r.width
}

// pass r by pointer so function can modify it
// note that double() is not required by shape interface
// but that doesn't matter
func (r *rect) double() {
	r.height *= 2
	r.width *= 2
}

func (t tri) area() float64 {
	return 1. / 2. * t.height * t.base
}

func (t tri) perimeter() float64 {
	// assumes isoleces triangle
	c := math.Sqrt(math.Pow(t.height, 2) + math.Pow(1./2.*t.base, 2))
	return t.base + 2*c
}

func (c circ) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circ) perimeter() float64 {
	// adding period to 2 to specify float
	return 2. * math.Pi * c.radius
}

func measure(s shape) {
	fmt.Printf("Shape: %s\n", s.name())
	fmt.Printf("Area: %f\n", s.area())
	fmt.Printf("Perimeter: %f\n\n", s.perimeter())
}

func main() {
	// make a slice of shapes of length 0
	shapes := make([]shape, 0)
	//shapes := []shape{} // alternate syntax using type deduction

	r1 := rect{"rectangle 1", 10, 5}
	shapes = append(shapes, r1)

	r2 := r1 // copy r1
	r2.double()
	r2.s = "rectangle 2"
	// modify member
	//measure(r1)
	//measure(r2)
	// r1 is unchanged
	shapes = append(shapes, r2)

	t1 := tri{"triangle 1", 10, 5}
	shapes = append(shapes, t1)

	c1 := circ{"circle 1", 1}
	shapes = append(shapes, c1)

	for _, s := range shapes {
		// measure takes argument of interface type shape
		// so works on rect, circ, tri
		measure(s)
	}
}
