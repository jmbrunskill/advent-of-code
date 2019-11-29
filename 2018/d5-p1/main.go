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
	//Print the result
	fmt.Println(len(react(string(b))))
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
