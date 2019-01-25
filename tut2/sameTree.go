package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

func Walker(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  Walker(t.Left, ch)
//  fmt.Println(t.Value)
  ch <- t.Value
  Walker(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  Walker(t, ch)
  close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  for v := range ch1 {
    if v != <-ch2 {
      return false
    }
  }
  _, ok := <- ch2
  if ok {
    return false
  }
  return true
}

func main() {
  fmt.Printf("Tree.New(1) == Tree.New(1) returns %v\n",
    Same(tree.New(1), tree.New(1)))
  fmt.Printf("Tree.New(1) == Tree.New(2) returns %v\n",
      Same(tree.New(1), tree.New(3)))
}
