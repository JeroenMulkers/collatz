/*
   Solves euler project problem 14
*/
package main

import (
	"fmt"
	"github.com/jeroenmulkers/collatz"
)

// returns the node with the highest order and with a value below 'maxValue'
func nodeWithHighestOrder(maxValue int) *collatz.Node {

	nodeFound := collatz.One

	for value := 1; value <= maxValue; value++ {
		node := collatz.NewNode(value)
		if node.Order() > nodeFound.Order() {
			nodeFound = node
		}
	}

	return nodeFound
}

func main() {
	N := 1000000
	fmt.Println(nodeWithHighestOrder(N).Value())
}
