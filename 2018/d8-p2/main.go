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
	nodeValue     int
	subnodes      []*node
}

func (n node) String() string {
	return fmt.Sprintf("{id: %c, ns: %d, md: %d, v: %d, sn: %v}", n.id, n.nodeCount, n.metaDataCount, n.nodeValue, n.subnodes)
}

func (n *node) addSubNode(sn *node) {
	n.subnodes = append(n.subnodes, sn)

}

type nodeStack []*node

func (s nodeStack) Push(v *node) nodeStack {
	return append(s, v)
}

func (s nodeStack) Pop() (nodeStack, *node) {
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

	var nodeID int = 'A'
	stack := make(nodeStack, 0)

	parserState := nodeCount //Expect node count first
	var head, current, parent *node

	nums := mapToInt(strings.Split(license, " "))
	for i, n := range nums {
		// if i > 0 {
		// 	fmt.Printf(",")
		// }
		//fmt.Printf("%d", n)
		fmt.Printf("%d:%d\n", i, n)

		if parserState == nodeCount {
			//New Node
			parent = current
			current = &node{id: nodeID, nodeCount: n}
			if head == nil {
				head = current
				parent = current
			} else {
				parent.addSubNode(current)
			}
			nodeID++
			// fmt.Printf("Started Reading Node %c, %d subnodes\n", current.id, n)
			parserState = metaDataCount // next number should be a metaDatacount
		} else if parserState == metaDataCount {
			//Read the metadataCount
			current.metaDataCount = n
			// fmt.Printf("Metadata Count for Node %c is %d\n", current.id, n)
			if current.nodeCount > 0 {
				stack = stack.Push(current)
				parserState = nodeCount //Next number should be a node count(starting new node)
				// fmt.Printf("More nodes expected for Node %c adding to stack\n", current.id)
			} else {
				//Last record in the stack so must be reading a metadatanode now
				parserState = metaDataValue
				// fmt.Printf("No more nodes expected for Node %c\n", current.id)
			}
		} else if parserState == metaDataValue {

			// fmt.Printf("metadata for Node %c - %d\n", current.id, n)

			if len(current.subnodes) == 0 && current.nodeCount == 0 {
				//This is  a leaf node - sum the metadata against this current node
				current.nodeValue += n
			} else {
				//This nodes' value needs to look at the sub nodes
				if n > 0 && n <= len(current.subnodes) {
					// fmt.Printf("Added value to node %c - sub node %d = %d\n", current.id, n, current.subnodes[n-1].nodeValue)
					current.nodeValue += current.subnodes[n-1].nodeValue
				} else {
					// fmt.Printf("Unable tl add the %dth subnode value to node %v\n", n, current)

				}
			}

			current.metaDataCount--

			if current.metaDataCount == 0 {
				// fmt.Printf("Finished metadata for Node %c value=%d\n", current.id, current.nodeValue)
				if len(stack) > 0 {
					stack, current = stack.Pop()
					current.nodeCount--
					if current.nodeCount > 0 {
						parserState = nodeCount
						stack = stack.Push(current)
					} else if current.metaDataCount > 0 {
						parserState = metaDataValue
					} else {
						parserState = done
					}
				} else {
					parserState = done
				}

				// fmt.Printf("Checking back to process Node %c - %d nodes to go\n", current.id, current.nodeCount)
			} else {
				// fmt.Printf("Expecting %d more metadata for Node %c value=%d\n", current.metaDataCount, current.id, current.nodeValue)
			}
		}
	}
	fmt.Println()

	result := head.nodeValue
	// fmt.Println(head)
	return fmt.Sprintf("%d", result)
}
