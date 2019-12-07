package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

type lineSegment struct {
	origin xy
	dest   xy
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//The manhattan Distance is the distance from the origin 0,0
func manhattanDistanceFromOrigin(x, y int) int {
	//eg. the total distance away from the origin in x plus the total distance away from the origin in y
	return abs(x) + abs(y)
}

//The manhattan Distance between points
func manhattanDistance(a, b xy) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func createLineSegment(origin xy, path string) lineSegment {
	ls := lineSegment{}

	// fmt.Println(path, "-", string(path[0]))
	ls.origin = origin

	direction := string(path[0])
	length, err := strconv.Atoi(string(path[1:]))
	if err != nil {
		panic("That aint an int!")
	}

	switch direction {
	case "R":
		//Right
		ls.dest = xy{origin.x + length, origin.y}
	case "L":
		//Left
		ls.dest = xy{origin.x - length, origin.y}
	case "U":
		//Up
		ls.dest = xy{origin.x, origin.y + length}
	case "D":
		//Down
		ls.dest = xy{origin.x, origin.y - length}
	}
	return ls
}

func lineSegments(wire []string) []lineSegment {
	origin := xy{0, 0}
	lineSegments := make([]lineSegment, 0)
	for _, path := range wire {
		ls := createLineSegment(origin, path)
		//New origin is the end of the new segment
		origin = ls.dest
		lineSegments = append(lineSegments, ls)
	}
	return lineSegments
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pointInSegment(p xy, line lineSegment) bool {
	//A Point is in a segment when
	// fmt.Println("pointInSegment", p, line)
	if p.x >= min(line.origin.x, line.dest.x) && p.x <= max(line.origin.x, line.dest.x) {
		// fmt.Println("OK IN X")
		if p.y >= min(line.origin.y, line.dest.y) && p.y <= max(line.origin.y, line.dest.y) {
			// fmt.Println("OK IN Y")
			return true
		}
	}

	return false
}

func lineIntersectionPoint(a, b lineSegment) *xy {

	//Two lines intersect if the first line is vertical origin x == dest x and the second is horizontal orgin
	//AND the point a.x, b.y is contained within both lines

	// fmt.Println("lineIntersectionPoint", a, b)

	if a.origin.x == a.dest.x {
		// fmt.Println("A is vertical")
		//A is Vertical
		if b.origin.y == b.dest.y {
			// fmt.Println("B is horizontal")
			//B is horizontal
			pnt := xy{a.origin.x, b.dest.y}
			if pointInSegment(pnt, a) && pointInSegment(pnt, b) {
				return &pnt
			}
		}
	} else if a.origin.y == a.dest.y {
		// fmt.Println("A is horizontal")
		//A is Hoizontal
		if b.origin.x == b.dest.x {
			// fmt.Println("B is vertical")
			//B is Vertical
			pnt := xy{b.origin.x, a.dest.y}
			if pointInSegment(pnt, a) && pointInSegment(pnt, b) {
				return &pnt
			}
		}
	} else {
		//Must be an error?
		panic(fmt.Sprintf("Diagonal Line %v", a))
	}

	//Return nil if there is no intersection
	return nil
}

func calcCrossDistance(wire1 []string, wire2 []string) int {

	//Here's the approach
	//1.create a list of line segments relative to origin
	//For each line segment check if there is an intersection, return this intersection point
	//Choose the lowest none 0 distance

	seg1 := lineSegments(wire1)
	seg2 := lineSegments(wire2)

	minDist := 999999
	for _, s1 := range seg1 {
		// fmt.Printf("Line 1 %v \n", s1)
		for _, s2 := range seg2 {
			// fmt.Printf("Line 2 %v \n", s2)
			p := lineIntersectionPoint(s1, s2)
			if p != nil {
				// fmt.Printf("Lines %v and %v intersect at %v\n", s1, s2, p)
				dist := manhattanDistanceFromOrigin(p.x, p.y)
				if dist > 0 && dist < minDist {
					minDist = dist
				}
			}
		}

	}

	return minDist
}

func delayFromOrigin(l1, l2 []lineSegment, p xy) int {
	// fmt.Println(l1, l2, p)
	totalDist := 0
	for i, l := range l1 {
		if i == len(l1)-1 {
			totalDist += manhattanDistance(l.origin, p)
		} else {
			totalDist += manhattanDistance(l.origin, l.dest)
		}
	}
	for i, l := range l2 {
		if i == len(l2)-1 {
			totalDist += manhattanDistance(l.origin, p)
		} else {
			totalDist += manhattanDistance(l.origin, l.dest)
		}
	}

	return totalDist
}

func calcCrossDelay(wire1 []string, wire2 []string) int {

	//Here's the approach
	//1.create a list of line segments relative to origin
	//For each line segment check if there is an intersection, return this intersection point
	//Choose the lowest none 0 distance

	seg1 := lineSegments(wire1)
	seg2 := lineSegments(wire2)

	minDist := 999999
	for i, s1 := range seg1 {
		// fmt.Printf("Line 1 %v \n", s1)
		for j, s2 := range seg2 {
			// fmt.Printf("Line 2 %v \n", s2)
			p := lineIntersectionPoint(s1, s2)
			if p != nil {
				// fmt.Printf("Lines %v and %v intersect at %v\n", s1, s2, p)
				dist := delayFromOrigin(seg1[0:i], seg2[0:j], *p)
				if dist > 0 && dist < minDist {
					minDist = dist
				}
			}
		}

	}

	return minDist
}

func processInput(f io.Reader) string {
	s := bufio.NewScanner(f)

	var lines [][]string //slice of slices of strings
	lines = make([][]string, 0, 0)

	for s.Scan() {
		// fmt.Println(s.Text())

		lines = append(lines, strings.Split(s.Text(), ","))

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	// fmt.Println("l1", lines[0])

	// fmt.Println("l2", lines[1])

	result := calcCrossDelay(lines[0], lines[1])

	return fmt.Sprintf("%v", result)
}
