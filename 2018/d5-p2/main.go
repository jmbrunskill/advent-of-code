package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Can't read input file", err)
	}

	polymer := string(b)

	minLength := len(polymer)
	// minPolymer := polymer
	for _, u := range reactionUnits() {
		fmt.Print(string(u))
		p := removeUnits(u, polymer)
		reactedP := react(p)
		if len(reactedP) < minLength {
			minLength = len(reactedP)
			// minPolymer = reactedP
		}
		fmt.Println("=", len(reactedP))
	}
	fmt.Println()

	fmt.Println(minLength)
}

func removeUnits(unit byte, polymer string) string {
	newPolymer := ""
	for i := 0; i < len(polymer); i++ {
		if polymer[i] == unit || doesReact(polymer[i], unit) {
			continue
		} else {
			newPolymer += string(polymer[i])
		}
	}
	return newPolymer
}

func reactionUnits() [26]byte {
	var unitList [26]byte
	//Assume only a-z reaction units
	for i := 0; i < 26; i++ {
		unitList[i] = 'a' + byte(i)
	}
	return unitList
}

func doesReact(a, b byte) bool {
	//thanks Fransec - https://github.com/campoy/advent-of-code-2018/blob/master/day05-p1/main.go
	const diff = 'a' - 'A'
	// return true
	return a+diff == b || b+diff == a
}

func react(polymer string) string {
	// fmt.Println("Reacting: ", polymer)
	i := 0
	for {
		//loop indefinitely

		if i < len(polymer)-1 {
			// fmt.Println(polymer[:i], "[", string(polymer[i]), string(polymer[i+1]), "]", polymer[i+2:])
			if doesReact(polymer[i], polymer[i+1]) {
				polymer = polymer[:i] + polymer[i+2:]
				if i > 0 {
					//go back to check if the previous char needs to be reacted
					i--
				}
			} else {
				i++
			}
		} else {
			//we got to the end of the polymer
			break
		}

	}
	// fmt.Println("Result: ", polymer)
	return polymer
}
