package main

type Bread struct {
  name string
  m map[string]Item
  weight float32
  baking Baking
}

type Baking struct {
  bakeTime, coolTime, temp int
}

type Item struct {
  weight int
}

type Baker interface {
  shoppingList(map[string]Item) (map[string]Item, map[string]Item)
  printBakeInstructions()
  printBreadInfo()
}
