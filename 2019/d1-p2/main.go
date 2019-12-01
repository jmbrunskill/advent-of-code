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

func calculateFuel(mass int) int {
	fuelSum := (mass / 3) - 2
	fuel := fuelSum
	for {
		fuel = (fuel / 3) - 2
		if fuel > 0 {
			fuelSum = fuelSum + fuel
		} else {
			break
		}
	}

	return fuelSum
}

func processInput(f io.Reader) string {
	sum := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		sum += calculateFuel(n)
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}
	return fmt.Sprintf("%d", sum)
}
