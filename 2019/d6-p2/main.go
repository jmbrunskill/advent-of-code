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

func orbitTransfers(orbitsReverse map[string]string, you, santa string) int {
	maxHops := 10000
	hopCount := 0
	youHops := make(map[string]int)

	//Find hops between YOU and COM
	p := you

	for i := 0; i < maxHops; i++ {
		// fmt.Printf("hop(%v), evaluating %v\n", you, p)
		youHops[p] = i
		if p == "COM" {
			break
		} else {
			hopCount++
			p = orbitsReverse[p]

		}
	}

	//Find hops from santa to a hops for you
	p = santa
	hopCount = 0
	for i := 0; i < maxHops; i++ {
		// fmt.Printf("hop(%v), evaluating %v\n", santa, p)
		if youHops[p] > 0 {
			return (youHops[p] - 1) + (hopCount - 1)
		}
		if p == "COM" {
			return -1
			break
		} else {
			hopCount++
			p = orbitsReverse[p]
		}
	}

	return -2
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

	orbitTransfers := orbitTransfers(orbitsReverse, "YOU", "SAN")

	return fmt.Sprintf("%d", orbitTransfers)
}
