package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func findLeaves(orbits, orbitsReverse map[string]string) []string {
	leaves := []string{}
	for k, _ := range orbitsReverse {
		fmt.Println(k, orbits[k])
		if orbits[k] == "" {
			leaves = append(leaves, k)
		}

	}
	return leaves
}

func orbitCountPlanet(orbitsReverse map[string]string, planet string) int {
	orbitCount := 0
	p := planet

	maxCount := 10000

	for i := 0; i < maxCount; i++ {
		// fmt.Printf("orbitCountPlanet(%v), evaluating %v\n", planet, p)
		if p == "COM" {
			return orbitCount
		} else {
			orbitCount++
			p = orbitsReverse[p]
		}
	}
	return orbitCount
}

func orbitCount(orbitsReverse map[string]string) int {
	orbitCount := 0
	calcd := make(map[string]bool)
	for k, _ := range orbitsReverse {
		if !calcd[k] {
			kCount := orbitCountPlanet(orbitsReverse, k)
			// fmt.Printf("Orbit Count for %v:%v\n", k, kCount)
			orbitCount += kCount
		}
	}
	return orbitCount
}

func processInput(f io.Reader) string {
	orbits := make(map[string]string)
	orbitsReverse := make(map[string]string)

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		obs := strings.Split(s.Text(), ")")

		if len(obs) < 2 {
			continue
		}
		orbits[obs[0]] = obs[1]
		orbitsReverse[obs[1]] = obs[0]
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	// fmt.Println(orbits)
	// fmt.Println(orbitsReverse)

	orbitCount := orbitCount(orbitsReverse)

	return fmt.Sprintf("%d", orbitCount)
}
