package main

import (
  "fmt"
)

type Person struct {
	Name string
}

type Saiyan struct {
  *Person
  Power int
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func (s *Saiyan) Exclaim() {
  fmt.Printf("My power is over %v!!!\n", s.Power)
}

// Write a function returning a new Saiyan structure.
// (factory pattern)
// return by pointer
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Person: &Person{name},
		Power: power,
	}
}

// Or return by value:
/*
func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}*/

// does not work as intended: s passed by copy
func Super(s Saiyan) {
  s.Power += 10000
}

func main() {
  // declare & initialize Saiyan
	goku := Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}
  // using factory
  trunks := NewSaiyan("Trunks", 500)
	goku.Introduce()
  trunks.Introduce()
  goku.Exclaim()
  Super(goku)
  goku.Exclaim()
}
