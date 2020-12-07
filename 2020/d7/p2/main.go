package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Bag struct {
	Name               string
	ContainedBagNames  []string
	ContainedBagCounts []int
}

var bagDefRegex = regexp.MustCompile("([a-z]+ [a-z]+) bags contain (.*)")
var bagContentsRegex = regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+) bag(s?)(.)?")

func parseBag(s string) *Bag {
	b := &Bag{}

	m := bagDefRegex.FindStringSubmatch(s)
	// fmt.Println(m)
	if m == nil {
		//didn't match the correct format
		return nil
	}

	b.Name = m[1]

	contents := strings.Split(m[2], ",")
	for _, containedBag := range contents {
		m := bagContentsRegex.FindStringSubmatch(containedBag)
		// fmt.Println(containedBag, " Contents Match", m)
		if m == nil {
			//didn't match the correct format
			return b
		}

		numBags, err := strconv.Atoi(m[1])
		if err != nil {
			return nil
		}
		bagColour := m[2]

		b.ContainedBagCounts = append(b.ContainedBagCounts, numBags)
		b.ContainedBagNames = append(b.ContainedBagNames, bagColour)
	}

	return b
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f))
}

func countSubBags(bagName string, bagRules map[string]Bag, bagContainsCount map[string]int) int {
	count := 0

	if _, ok := bagContainsCount[bagName]; ok {
		//We've already calculated the number of bags contained here
		return bagContainsCount[bagName]
	}

	//We don't know how many bags we're going to get here so check all the contained bags (recursively)
	if bag, ok := bagRules[bagName]; ok {
		for i, name := range bag.ContainedBagNames {
			count += (bag.ContainedBagCounts[i] * (1 + countSubBags(name, bagRules, bagContainsCount)))
		}
	}

	//Record this result so we don't need to re-calculate
	bagContainsCount[bagName] = count

	return count
}

func processInput(f io.Reader) string {
	startTime := time.Now().Unix()
	result := 0

	bagRules := make(map[string]Bag)
	bagContainsCount := make(map[string]int)
	b := &Bag{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		b = parseBag(s.Text())
		bagRules[b.Name] = *b
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	for k, _ := range bagRules {
		count := countSubBags(k, bagRules, bagContainsCount)
		// fmt.Println(k, count)
		if count > 0 {
			result++
		}
	}

	result = bagContainsCount["shiny gold"]

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result %v in %d seconds\n", result, endTime-startTime)

	return fmt.Sprintf("%d", result)
}
