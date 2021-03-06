package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Can't read input file", err)
	}

	//Print the result
	fmt.Println(processInput(string(b)))
}

func mapToInt(vs []string) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		iv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("could not parse int %q: %v", v, err)
		}
		vsm[i] = iv
	}
	return vsm
}

type node struct {
	id            int
	nodeCount     int
	metaDataCount int
}

func (n node) String() string {
	return fmt.Sprintf("{id: %c, ns: %d, md: %d}", n.id, n.nodeCount, n.metaDataCount)
}

type nodeStack []node

func (s nodeStack) Push(v node) nodeStack {
	return append(s, v)
}

func (s nodeStack) Pop() (nodeStack, node) {
	// fmt.Println("POP", s)
	l := len(s)

	if len(s) < 1 {
		return make(nodeStack, 0), s[1]
	}
	return s[:l-1], s[l-1]
}

type parsedType byte

const (
	nodeCount parsedType = iota
	metaDataCount
	metaDataValue
	done
)

func processInput(license string) string {
	metadataSum := 0
	var nodeID int = 'A'
	stack := make(nodeStack, 0)

	parserState := nodeCount //Expect node count first
	currentNode := node{}

	nums := mapToInt(strings.Split(license, " "))
	for i, n := range nums {
		// if i > 0 {
		// 	fmt.Printf(",")
		// }
		//fmt.Printf("%d", n)
		fmt.Printf("%d:%d\n", i, n)

		if parserState == nodeCount {
			//New Node
			currentNode = node{id: nodeID, nodeCount: n}
			nodeID++
			// fmt.Printf("Started Reading Node %c, %d subnodes\n", currentNode.id, n)
			parserState = metaDataCount // next number should be a metaDatacount
		} else if parserState == metaDataCount {
			//Read the metadataCount
			currentNode.metaDataCount = n
			// fmt.Printf("Metadata Count for Node %c is %d\n", currentNode.id, n)
			if currentNode.nodeCount > 0 {
				stack = stack.Push(currentNode)
				parserState = nodeCount //Next number should be a node count(starting new node)
				// fmt.Printf("More nodes expected for Node %c adding to stack\n", currentNode.id)
			} else {
				//Last record in the stack so must be reading a metadatanode now
				parserState = metaDataValue
				// fmt.Printf("No more nodes expected for Node %c\n", currentNode.id)
			}
		} else if parserState == metaDataValue {

			// fmt.Printf("metadata for Node %c - %d\n", currentNode.id, n)
			metadataSum += n //add to the metadatasum
			currentNode.metaDataCount--

			if currentNode.metaDataCount == 0 {
				// fmt.Printf("Finished metadata for Node %c\n", currentNode.id)
				if len(stack) > 0 {

					stack, currentNode = stack.Pop()
					currentNode.nodeCount--
					if currentNode.nodeCount > 0 {
						parserState = nodeCount
						stack = stack.Push(currentNode)
					} else if currentNode.metaDataCount > 0 {
						parserState = metaDataValue
					} else {
						parserState = done
					}
				} else {
					parserState = done
				}

				// fmt.Printf("Checking back to process Node %c - %d nodes to go\n", currentNode.id, currentNode.nodeCount)

			}
		}
	}
	fmt.Println()

	return fmt.Sprintf("%d", metadataSum)
}
