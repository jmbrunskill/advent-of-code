package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

type reaction struct {
	name        string
	outputCount int
	inputNames  []string
	inputCounts []int
}

func parseReaction(s string) reaction {
	r := reaction{}
	var err error

	iAndO := strings.Split(s, "=>")
	if len(iAndO) < 1 {
		panic(fmt.Sprintf("INVALID REACTION %v", s))
	}

	inputs := strings.Split(iAndO[0], ",")

	outputInfo := strings.Fields(iAndO[1])
	r.name = outputInfo[1]
	r.outputCount, err = strconv.Atoi(outputInfo[0])
	if err != nil {
		panic("Invalid Output Count")
	}

	for _, input := range inputs {

		inputInfo := strings.Fields(input)
		name := inputInfo[1]
		inputCount, err := strconv.Atoi(inputInfo[0])
		if err != nil {
			panic("Invalid Input Count")
		}
		r.inputCounts = append(r.inputCounts, inputCount)
		r.inputNames = append(r.inputNames, name)
	}

	return r

}

type reactionLab struct {
	chemCount   map[string]int
	chemList    map[string]reaction
	oreConsumed int
}

func (lab reactionLab) String() string {
	str := "+---------------------------------------------+\n"
	var keys []string
	for k := range lab.chemCount {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		str += fmt.Sprintln(k, lab.chemCount[k])

	}
	str += fmt.Sprintln("Current Ore Consumed", lab.oreConsumed)
	str += "+---------------------------------------------+\n"

	return str
}

func (lab *reactionLab) useOrMake(chem string, count int) {
	// fmt.Printf("useOrMake(%s,%d)\n", chem, count)

	//Special Case for ORE
	if chem == "ORE" {
		lab.oreConsumed += count
		return
	}
	// fmt.Printf("useOrMake(%s,%d) - We have %v available\n", chem, count, lab.chemCount[chem])

	if lab.chemCount[chem] < count {
		//We don't have enough of this chemical, we'll need to make some more
		amountNeeded := count - lab.chemCount[chem]
		amountToProduce := 0
		reactionsNeeded := 0
		for amountToProduce < amountNeeded {
			amountToProduce += lab.chemList[chem].outputCount
			reactionsNeeded++
		}
		// fmt.Printf("We need to run this reaction %d times\n", reactionsNeeded)

		//Otherwise, consume or use all the other ingredients needed
		for i := 0; i < len(lab.chemList[chem].inputNames); i++ {
			//Produce reactionsNeeded of this chemical
			lab.useOrMake(lab.chemList[chem].inputNames[i], lab.chemList[chem].inputCounts[i]*reactionsNeeded)
		}
		lab.chemCount[chem] += amountToProduce
	}

	//We already have enough of this chemical, just use it...
	lab.chemCount[chem] -= count
	// fmt.Println(lab)
}

func processInput(f io.Reader) string {
	lab := reactionLab{}
	lab.chemCount = make(map[string]int)
	lab.chemList = make(map[string]reaction)

	//Load the Chemical Processes
	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		r := parseReaction(s.Text())
		if lab.chemCount[r.name] == 0 {
			lab.chemCount[r.name] = 0
		} else {
			fmt.Printf("WARNING DUPLICATE WAY TO MAKE %v (%v) and (%v)\n", r.name, lab.chemCount[r.name], s.Text())
		}
		lab.chemList[r.name] = r

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	lab.useOrMake("FUEL", 1)
	fmt.Println(lab)

	return fmt.Sprintf("%d", lab.oreConsumed)
}
