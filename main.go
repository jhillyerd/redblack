package main

import (
	"fmt"
)

func main() {
	tree := &RBNode{
		1,
		false,
		&RBNode{2, true, nil, nil},
		&RBNode{3, true, nil, nil},
	}
	fmt.Println(tree)
}
