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

type xy struct {
	x int
	y int
}

type fabric struct {
	points   map[xy]int
	validIds map[int]bool
}

func (f *fabric) claim(id, x, y, w, h int) {

	//Assume this id is valid to start with
	f.validIds[id] = true

	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			// fmt.Println("Claiming", i, j)
			point := xy{i, j}
			if f.points[point] > 0 || f.points[point] == -1 {
				// fmt.Println("duplicateClaim", i, j, f.points[point])

				//Mark this ID as invalid
				f.validIds[id] = false
				if f.points[point] != -1 {
					//Also invalidate the id using this space
					f.validIds[f.points[point]] = false
				}

				f.points[point] = -1
			} else {
				f.points[point] = id
			}
		}

	}
}

func (f *fabric) countDuplicated() int {
	count := 0
	for p := range f.points {
		if f.points[p] == -1 {
			count++
		}
	}
	return count
}

func (f *fabric) findValidId() int {
	for k, v := range f.validIds {
		if v {
			return k
		}
	}
	return -1
}

func (f fabric) Print() {
	var maxX, maxY int
	for p := range f.points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	//Print Y Headers

	for y := 0; y <= maxY; y++ {
		if y > 0 {
			fmt.Printf(",%d", y)
		} else {
			fmt.Printf(" :%d", y)
		}
	}
	fmt.Println()

	for y := 0; y <= maxY; y++ {
		fmt.Printf("%d:", y)
		for x := 0; x <= maxX; x++ {
			if x > 0 {
				fmt.Printf(",%d", f.points[xy{x, y}])
			} else {
				fmt.Printf("%d", f.points[xy{x, y}])
			}

		}
		fmt.Println()
	}

}

func processInput(f io.Reader) string {

	var fab fabric
	fab.points = make(map[xy]int)
	fab.validIds = make(map[int]bool)

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatalf("could not decode line %s: %v", s.Text(), err)
		}
		// fmt.Printf("#%d @ %d,%d: %dx%d\n", id, x, y, w, h)
		fab.claim(id, x, y, w, h)
		// fab.Print()
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}
	// fab.Print()
	id := fab.findValidId()
	return fmt.Sprintf("%d", id)

}
