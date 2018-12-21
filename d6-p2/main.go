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
	fmt.Println(processInput(f, 10000))
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

func (p *plane) calcArea(size int) int {

	area := 0

	for y := 0; y <= p.maxY; y++ {
		for x := 0; x <= p.maxX; x++ {
			pnt := xy{x, y}
			if p.distSum(pnt) < size {
				area++
			}

		}
	}
	return area
}

func (p *plane) distSum(pnt xy) int {
	distSum := 0

	for loc := range p.points {
		dist := dist(loc, pnt)
		distSum += dist
	}
	return distSum
}

func (p plane) Print(size int) {
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

			if (p.distSum(xy{x, y}) < size) {
				color.Set(color.BgGreen)
				color.Set(color.FgWhite)
			} else {
				color.Set(color.BgRed)
				color.Set(color.FgWhite)
			}

			fmt.Printf("%02d", p.closest[xy{x, y}])

			color.Unset()

		}
		fmt.Println()
	}

}

func processInput(f io.Reader, size int) string {
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

	area := p.calcArea(size)
	// p.Print(size)
	return fmt.Sprintf("%d", area)
}
