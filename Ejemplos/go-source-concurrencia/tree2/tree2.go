package main

import (
	"fmt"
	"tree2/tree"
)

func lt(x, y int) bool {
	return x <= y
}

func main() {
	var t *tree.Tree[int]
	for _, i := range []int{50, 30, 90, 40, 60, 10, 80, 35, 55} {
		t = t.Insert(i, lt)
	}
	fmt.Println("Tree:", t.GetAll())

	fmt.Println("Paths:")
	allPaths := t.AllPaths()
	for _, path := range allPaths {
		fmt.Println(path)
	}
}
