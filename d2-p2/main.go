package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f))
}

func processInput(f io.Reader) string {

	var suffixTree SuffixTree

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println("\n***PROCESSING****\n", s.Text())
		str, found := suffixTree.findOrAddString(s.Text())
		if found {
			return str
		}

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	return "No Matches Found"

}

// SuffixTree implements a special purpose suffix tree, returns string when single character miss match found
// Assumes all inputs are the same length
type SuffixTree struct {
	m map[byte]*SuffixTreeNode
}

type SuffixTreeNode struct {
	path string
	m    map[byte]*SuffixTreeNode
}

func (t *SuffixTree) findOrAddString(str string) (string, bool) {
	thisChar := str[0]
	thisSuffix := str[1:]

	if t.m == nil {
		t.m = make(map[byte]*SuffixTreeNode)
	}

	if t.m[thisChar] == nil {
		// fmt.Printf("findOrAddString(%v) - adding %s => %v \n", str, string(thisChar), thisSuffix)
		t.m[thisChar] = &SuffixTreeNode{path: string(thisChar)}
		// fmt.Printf("Added N:%v\n", t.m[thisChar].path)
	}
	return t.m[thisChar].findOrAddString(thisSuffix, string(thisChar))
}

func (t *SuffixTreeNode) findOrAddString(str, matched string) (string, bool) {

	// fmt.Printf("N:%v - findOrAddString(%v,%v)\n", t.path, str, matched)
	if len(str) == 0 {
		// fmt.Printf("N:%v - Nothing to Add", t.path)
		return "Nothing to Add", false
	}

	myFirstChar := str[0]
	mySuffix := str[1:]

	// fmt.Printf("me(%s,%s)\n", string(myFirstChar), mySuffix)

	//make map if not already created
	if t.m == nil {
		// fmt.Printf("N:%v - making map\n", t.path)
		t.m = make(map[byte]*SuffixTreeNode)
	}

	if t.suffixRangeMatch(myFirstChar, mySuffix) {
		//We found a match!
		return (matched + mySuffix), true
	}

	//No match, so add to the tree
	if t.m[myFirstChar] == nil {
		t.m[myFirstChar] = &SuffixTreeNode{path: (t.path + string(myFirstChar))}
		// fmt.Printf("Added N:%s to %s\n", t.m[myFirstChar].path, string(myFirstChar))
	}

	return t.m[myFirstChar].findOrAddString(mySuffix, (matched + string(myFirstChar)))

}

//This function checks for exact match one character downstream
func (t *SuffixTreeNode) suffixRangeMatch(currentChar byte, str string) bool {
	if len(str) == 0 {
		// fmt.Printf("N:%v - suffixRangeMatch(%v) - empty string is not a match\n", t.path, str)
		return false
	}

	for k, v := range t.m {
		if v != nil {
			if k == currentChar {
				// fmt.Printf("skipping %s\n", string(k))
				continue
			}
			// fmt.Printf("N:%v - suffixRangeMatch(%v) checking for match under %v\n", t.path, str, k)
			if v.suffixMatch(str) {
				return true
			}
		}

	}
	return false
}

//Checks to see if suffix is in the tree
func (t *SuffixTreeNode) suffixMatch(str string) bool {
	// fmt.Printf("N:%v - suffixMatch(%v)\n", t.path, str)

	if len(str) == 0 {
		// fmt.Printf("N:%v -  suffixMatch(%v) - empty string so must match\n", t.path, str)
		return true
	}
	myFirstChar := str[0]

	if t.m[myFirstChar] == nil {
		//suffix not in the tree
		// fmt.Printf("N:%v - suffixMatch(%v) - not in tree %v\n", t.path, str, myFirstChar)
		return false
	}

	return t.m[myFirstChar].suffixMatch(str[1:])
}
