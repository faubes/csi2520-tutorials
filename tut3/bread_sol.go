package main

import (
	"fmt"
)

type Item struct {
	weight int
}

type Baking struct {
	bakeTime, coolTime, temperature int
}

type Bread struct {
	name         string
	ingredients  map[string]Item
	instructions Baking
	weight       float32
}

type BakedGoods interface {
	shoppingList(map[string]Item) (map[string]Item, map[string]Item)
	printBakeInstructions()
	printBreadInfo()
}

func NewBread() (b *Bread) {
	b = new(Bread)
	b.name = "Whole Wheat"
	b.ingredients = map[string]Item{
		"whole wheat flour": Item{500},
		"yeast":             Item{25},
		"salt":              Item{25},
		"sugar":             Item{50},
		"butter":            Item{50},
		"water":             Item{350},
	}
	b.instructions = Baking{120, 60, 180}
	b.weight = b.sumWeights()
	return
}

func NewBreadVariation(name string, add, remove map[string]Item) (b *Bread) {
	b = NewBread()
	b.name = name
	for nm, it := range add {
		if _, ok := b.ingredients[nm]; ok {
			b.ingredients[nm] = Item{b.ingredients[nm].weight + it.weight}
		} else {
			b.ingredients[nm] = Item{it.weight}
		}
	}
	for nm, it := range remove {
		if _, ok := b.ingredients[nm]; ok {
			b.ingredients[nm] = Item{b.ingredients[nm].weight - it.weight}
			if b.ingredients[nm].weight <= 0 {
				delete(b.ingredients, nm)
			}
		}
	}
	b.weight = b.sumWeights()
	return
}

func (b *Bread) sumWeights() (weight float32) {
	for _, it := range b.ingredients {
		weight += float32(it.weight)
	}
	weight /= 1000.0
	return
}

func (b *Bread) shoppingList(have map[string]Item) (need, remain map[string]Item) {
	need = make(map[string]Item)
	remain = make(map[string]Item)
	for k, v := range have {
		remain[k] = v
	}
	for nm, it := range b.ingredients {
		delta := it.weight
		if _, ok := remain[nm]; ok {
			delta = it.weight - remain[nm].weight
			if delta < 0 {
				remain[nm] = Item{-delta}
				delta = 0
			}
		}
		if delta > 0 {
			need[nm] = Item{delta}
		}
	}
	return
}

func (b *Bread) printBakeInstructions() {
	fmt.Printf("Bake at %d Celsius for %d minutes and let cool for %d minutes.\n",
		b.instructions.temperature, b.instructions.bakeTime, b.instructions.coolTime)
}


func (b *Bread) printBreadInfo() {
	fmt.Printf("%s bread \n", b.name)
	fmt.Println(b.ingredients)
	fmt.Printf("Weight %.3f kg\n\n", b.weight)
	return
}


func addItemMap( mA, mB map[string]Item) (mC map[string]Item) {
	mC = make(map[string]Item)
	for k, v := range mA {
		mC[k] = v
	}
	for nm, it := range mB {
		if it2, ok := mC[nm]; ok {
			it2.weight += it.weight
			mC[nm] = it2
		} else {
			mC[nm] = mB[nm]
		}
	}
	return
}


func main() {
	breads := []BakedGoods{NewBread(),
		NewBreadVariation("Sesame", map[string]Item{"white flour": Item{200}, "sesame": Item{50}}, map[string]Item{"whole wheat flour": Item{250}})}
	for _, val := range breads {
		val.printBreadInfo()
	}
	fmt.Println()
	have := map[string]Item{"whole wheat flour": Item{5000}, "salt": Item{500}, "sugar": Item{1000}}
	need := make(map[string]Item)
	for _, val := range breads {
		var needB map[string]Item
		needB, have = val.shoppingList(have)
		need = addItemMap( need, needB )
	}
	fmt.Println("Shopping List:")
	fmt.Println(need)
	fmt.Println()
	fmt.Println("Baking Instructions:")
	for _, val := range breads {
		val.printBakeInstructions()
	}
	return
}
