package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(a, b xy) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

type plane struct {
	maxX    int
	maxY    int
	points  map[xy]int
	closest map[xy]int
	areas   map[xy]int
}

func (p *plane) addPoint(id, x, y int) {

	if p.points == nil {
		p.points = make(map[xy]int)
		p.closest = make(map[xy]int)
	}

	//Capture min & max X to give bounds to our searching later
	if p.maxX < x {
		p.maxX = x
	}
	if p.maxY < y {
		p.maxY = y
	}

	p.points[xy{x, y}] = id
}

func (p *plane) calcAll() int {
	p.areas = make(map[xy]int)

	maxArea := 0

	for y := 0; y <= p.maxY; y++ {
		for x := 0; x <= p.maxX; x++ {
			pnt := xy{x, y}
			id, loc := p.closestTo(pnt)
			if id > 0 {
				p.closest[pnt] = id
				p.areas[loc]++
				// fmt.Printf("Areas for %d is now %d\n", p.points[loc], p.areas[loc])
				if p.isBounded(loc) {
					if (p.areas[loc]) > maxArea {
						maxArea = p.areas[loc]
					}
				}
			}

		}
	}
	return maxArea
}

func (p *plane) closestTo(pnt xy) (int, xy) {
	minDist := -1
	minLocID := 0
	minLocPnt := xy{}

	for loc, id := range p.points {
		dist := dist(loc, pnt)
		if minDist == -1 || dist < minDist {
			minDist = dist
			minLocID = id
			minLocPnt = loc
		} else if dist == minDist {
			minDist = dist
			minLocID = -1
		}
	}
	return minLocID, minLocPnt
}

// p is surrounded if there are 4 points a, b, c, and d for which
// - a.x > p.x && a.y > p.y
// - b.x < p.x && b.y < p.y
// - c.x > p.x && c.y < p.y
// - d.x < p.x && d.y > p.y
func (p *plane) isBounded(pnt xy) bool {
	nw := false //North west point
	ne := false
	sw := false
	se := false
	for loc := range p.points {
		switch {
		case pnt.x > loc.x && pnt.y > loc.y:
			nw = true
		case pnt.x > loc.x && pnt.y < loc.y:
			ne = true
		case pnt.x < loc.x && pnt.y > loc.y:
			sw = true
		case pnt.x < loc.x && pnt.y < loc.y:
			se = true
		}

		if nw && ne && sw && se {
			// fmt.Printf("Point %d - %d,%d is bounded\n", p.points[pnt], pnt.x, pnt.y)
			return true //Early return
		}
	}
	// fmt.Printf("Point %d - %d,%d is NOT bounded\n", p.points[pnt], pnt.x, pnt.y)
	return false
}

func (p plane) Print() {
	//Print Headers
	for y := 0; y <= p.maxY; y++ {
		if y > 0 {
			fmt.Printf(",%d", y)
		} else {
			fmt.Printf(" :%d", y)
		}
	}
	fmt.Println()

	for y := 0; y <= p.maxY; y++ {
		fmt.Printf("%d:", y)
		for x := 0; x <= p.maxX; x++ {

			if x > 0 {
				fmt.Printf(",")
			}

			if (p.points[xy{x, y}] > 0) {
				color.Set(color.BgRed)
				color.Set(color.FgWhite)
			} else {
				color.Set(color.BgCyan)
				color.Set(color.FgRed)
			}

			fmt.Printf("%02d", p.closest[xy{x, y}])

			color.Unset()

		}
		fmt.Println()
	}

}

func processInput(f io.Reader) string {
	lineID := 1

	var p plane

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		var x, y int
		_, err := fmt.Sscanf(s.Text(), "%d, %d", &x, &y)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		p.addPoint(lineID, x, y)
		lineID++
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	maxArea := p.calcAll()
	// p.Print()
	return fmt.Sprintf("%d", maxArea)
}
