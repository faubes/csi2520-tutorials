package main

import (
  "fmt"
)

type Item struct {weight int}

type IngredientList map[string]Item

type Bread struct {
  name string
  ingredients IngredientList
  weight float32
  instructions Baking
}

type Baking struct {
  bakeTime, coolTime, temp int
}

type Baker interface {
  shoppingList(map[string]Item) (IngredientList, IngredientList)
  printBakeInstructions()
  printBreadInfo()
}

func NewBread() (*Bread) {
  b := &Bread{
    name: "Whole Wheat",
    ingredients: IngredientList{
      "whole wheat flour":{500},
      "yeast":{25},
      "salt":{25},
      "sugar":{50},
      "butter":{50},
      "water":{350},
    },
    instructions: Baking{
      bakeTime: 120,
      coolTime: 60,
      temp: 180,
    },
  }
  for _, v := range b.ingredients {
    b.weight += float32(v.weight)
  }
  return b
}

func main() {
  ww := NewBread()
  fmt.Printf("%v\n", ww)
}
