package main

import (
	"fmt"
	"publicChain/part1-Basic-Prototype/BLC"
)

func main() {

	block := BLC.NewBlock("this is block", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block)
}
