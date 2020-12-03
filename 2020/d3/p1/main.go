package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f, 1, 3))
}

func treesOnSlope(down, right int, treeline [][]bool) int {

	rowCount := len(treeline)
	colCount := len(treeline[0])
	// fmt.Printf("Found tree line %d by %d\n", rowCount, colCount)
	d := 0 //How far down we are
	r := 0 //How far right we are

	treeCount := 0

	for {
		//If we have got to the end of the tree line, retun
		if d >= rowCount {
			break
		}

		//If we've hit a tree, count it
		if treeline[d][r] {
			treeCount++
		}
		// fmt.Println(d, r, treeline[d][r])

		//Update our position
		d += down
		r += right
		r = r % (colCount)

	}

	return treeCount
}

func processInput(f io.Reader, down, right int) string {
	startTime := time.Now().Unix()
	result := 0

	treeline := make([][]bool, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		rowline := make([]bool, 0)
		for _, ch := range s.Text() {
			rowline = append(rowline, (ch == '#'))
		}
		treeline = append(treeline, rowline)
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	result = treesOnSlope(down, right, treeline)

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result in %d seconds\n", endTime-startTime)

	return fmt.Sprintf("%d", result)
}
