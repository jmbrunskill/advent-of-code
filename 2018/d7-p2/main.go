package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f, 5))
}

func processInput(f io.Reader, numWorkers int) string {
	var t tree
	//Init (Cheating as this should be an init function)
	t.nodes = make(map[byte]*node)
	t.workers = make(map[int]byte)
	for i := 0; i < numWorkers; i++ {
		t.workers[i] = ' '
	}

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

	for !t.inc() {
		t.Print()
	}

	return fmt.Sprintf("%d", t.second-1) //There's an extra increment in t.inc()
}

type node struct {
	name         byte
	time         int
	inprogress   bool
	Predecessors []byte
}

type tree struct {
	nodes   map[byte]*node
	workers map[int]byte
	second  int
}

func (t *tree) Print() {
	fmt.Printf("%d:", t.second)
	for _, v := range t.workers {
		if v != ' ' {

			fmt.Printf("%c (%d)", v, t.nodes[v].time)
		} else {
			fmt.Printf(" * ")
		}
	}
	fmt.Println()
}

func (t *tree) add(succ, pred byte) {
	const diff byte = 'A'
	if t.nodes[succ] == nil {
		t.nodes[succ] = &node{name: succ, time: int(succ-diff) + 60}
		fmt.Printf("Added %c %d\n", succ, t.nodes[succ].time)
	}
	if t.nodes[pred] == nil {
		t.nodes[pred] = &node{name: pred, time: int(pred-diff) + 60}
		fmt.Printf("Added %c %d\n", pred, t.nodes[pred].time)
	}
	t.nodes[succ].Predecessors = append(t.nodes[succ].Predecessors, pred)

}

func (t *tree) inc() bool {
	// fmt.Println("INC", t.second)
	//update all the active nodes...
	for i, w := range t.workers {
		// fmt.Printf("Worker %d doing %c\n", i, w)
		if w == ' ' {
			//This worker is doing nothing!
			continue
		} else if t.nodes[w].time > 0 {
			t.nodes[w].time--
		} else {
			t.complete(w)
			t.workers[i] = ' '
		}
	}

	//put idle hands to work
	for _, n := range t.readyNodes() {
		for i, w := range t.workers {
			// fmt.Printf("%c - %d\n", n.name, len(n.Predecessors))
			if w == ' ' {
				//put this worker to work
				t.nodes[n].inprogress = true
				t.workers[i] = n
				break
			}
		}
	}

	t.second++
	return len(t.nodes) == 0
}

func (t *tree) readyNodes() []byte {
	var readies []byte

	for _, n := range t.nodes {
		// fmt.Printf("%c - %d\n", n.name, len(n.Predecessors))
		if t.ready(n.name) {
			fmt.Printf("%c READY\n", n.name)
			readies = append(readies, n.name)
		}
	}
	sort.Slice(readies, func(i, j int) bool {
		return readies[i] < readies[j]
	})

	return readies
}

func (t *tree) ready(name byte) bool {
	if t.nodes[name].inprogress {
		return false
	}
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
