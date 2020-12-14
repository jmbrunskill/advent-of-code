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
	fmt.Println(processInput(f))
}

func countOccupied(seating [][]int) int {
	sum := 0
	for i := 0; i < len(seating); i++ {
		for j := 0; j < len(seating[i]); j++ {
			if seating[i][j] == 1 {
				sum++
			}
		}
	}
	return sum
}

func occupiedAdjacent(seating [][]int, i, j int) int {
	numOccupied := 0
	if i > 0 && j > 0 && seating[i-1][j-1] > 0 {
		numOccupied++
	}
	if i > 0 && seating[i-1][j] > 0 {
		numOccupied++
	}
	if i > 0 && j+1 < len(seating[i-1]) && seating[i-1][j+1] > 0 {
		numOccupied++
	}

	if j > 0 && seating[i][j-1] > 0 {
		numOccupied++
	}
	if j+1 < len(seating[i]) && seating[i][j+1] > 0 {
		numOccupied++
	}

	if i+1 < len(seating) && j > 0 && seating[i+1][j-1] > 0 {
		numOccupied++
	}
	if i+1 < len(seating) && seating[i+1][j] > 0 {
		numOccupied++
	}
	if i+1 < len(seating) && j+1 < len(seating[i+1]) && seating[i+1][j+1] > 0 {
		numOccupied++
	}
	return numOccupied
}

func updateSeats(seating [][]int) int {
	updatedSeats := 0

	//need a shadow copy of the seats as otherwise our incremental changes to the array will affect the results
	var shadowSeats [][]int
	for i := 0; i < len(seating); i++ {
		row := append([]int{}, seating[i]...)
		shadowSeats = append(shadowSeats, row)
	}

	for i := 0; i < len(seating); i++ {
		for j := 0; j < len(seating[i]); j++ {
			// fmt.Printf("i:%d,j:%d, occ: %d\n", i, j, occupiedAdjacent(shadowSeats, i, j))
			if seating[i][j] == 0 {
				if occupiedAdjacent(shadowSeats, i, j) == 0 {
					// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
					seating[i][j] = 1
					updatedSeats++
				}
			} else if seating[i][j] == 1 {
				if occupiedAdjacent(shadowSeats, i, j) >= 4 {
					// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
					seating[i][j] = 0
					updatedSeats++
				}
			}
		}
	}
	return updatedSeats
}

func printSeating(seating [][]int) {

	for x := 0; x < len(seating); x++ {
		for y := 0; y < len(seating[x]); y++ {
			switch seating[x][y] {
			case 0:
				fmt.Printf("L")
			case 1:
				fmt.Printf("#")
			case -1:
				fmt.Printf(".")
			default:
				fmt.Printf("?")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func processInput(f io.Reader) string {
	startTime := time.Now()
	result := 0

	seating := make([][]int, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		row := make([]int, 0)
		for _, ch := range s.Text() {
			if ch == 'L' {
				row = append(row, 0)
			} else if ch == '.' {
				row = append(row, -1)
			}

		}
		seating = append(seating, row)
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	updatedSeats := 1
	for {
		updatedSeats = updateSeats(seating)
		// printSeating(seating)
		fmt.Printf("Updated %d seats\n", updatedSeats)
		if updatedSeats == 0 {
			break
		}
	}

	result = countOccupied(seating)

	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
