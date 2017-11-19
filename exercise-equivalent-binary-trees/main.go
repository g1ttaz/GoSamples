package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk0 is the inner walk
func Walk0(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk0(t.Left, ch)
	ch <- t.Value
	Walk0(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walk0(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		i1, ok1 := <-ch1
		i2, ok2 := <-ch2
		if ok1 != ok2 || i1 != i2 {
			return false
		}
		if !ok1 || !ok2 {
			return true
		}
	}
}

func main() {
	ok := Same(tree.New(1), tree.New(1))
	fmt.Println("ok=", ok)
	nok := Same(tree.New(1), tree.New(2))
	fmt.Println("nok=", nok)
}
