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
	var t tree
	t.nodes = make(map[byte]*node)

	order := ""
	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		var pred, succ byte
		_, err := fmt.Sscanf(s.Text(), "Step %c must be finished before step %c can begin.", &pred, &succ)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		// fmt.Printf("%c then %c\n", pred, succ)

		t.add(succ, pred)

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	for {
		char, done := t.step()
		order += fmt.Sprintf("%c", char)
		if done {
			break
		}

	}

	return fmt.Sprintf("%s", order)
}

type node struct {
	name         byte
	Predecessors []byte
}

type tree struct {
	nodes map[byte]*node
}

func (t *tree) add(succ, pred byte) {
	if t.nodes[succ] == nil {
		t.nodes[succ] = &node{name: succ}
	}
	if t.nodes[pred] == nil {
		t.nodes[pred] = &node{name: pred}
	}
	t.nodes[succ].Predecessors = append(t.nodes[succ].Predecessors, pred)
}

func (t *tree) step() (byte, bool) {
	var nextNode byte
	nextNode = 'Z' //Assume we keep to ascii chars

	for _, n := range t.nodes {
		// fmt.Printf("%c - %d\n", n.name, len(n.Predecessors))
		if t.ready(n.name) {
			// fmt.Printf("%c READY\n", n.name)
			if n.name < nextNode {
				nextNode = n.name
			}
		}
	}

	//TODO remove this node, from predecessors list
	t.complete(nextNode)

	return nextNode, len(t.nodes) == 0
}

func (t *tree) ready(name byte) bool {
	if len(t.nodes[name].Predecessors) == 0 {
		return true
	}

	return false
}

func (t *tree) complete(name byte) {
	//remove from any predecessors
	for _, n := range t.nodes {
		for i, c := range n.Predecessors {
			if c == name {
				//remove element i from slice
				n.Predecessors = append(n.Predecessors[:i], n.Predecessors[i+1:]...)
			}
		}
	}
	//remove from map
	delete(t.nodes, name)
}
