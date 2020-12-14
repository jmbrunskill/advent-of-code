package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
func processInput(f io.Reader) string {
	startTime := time.Now()
	result := 0

	s := bufio.NewScanner(f)

	//Get Earliest Time
	s.Scan()
	earliestTime, err := strconv.Atoi(s.Text())
	if err != nil {
		return "ERROR NO EARLIEST TIME"
	}

	//Get schedules
	s.Scan()
	busStrings := strings.Split(s.Text(), ",")

	busNums := []int{}
	for _, bus := range busStrings {
		busid, err := strconv.Atoi(bus)
		if err == nil {
			busNums = append(busNums, busid)
		}

	}

	minLeaveTime := earliestTime * 2
	minLeaveBus := 0
	for _, busid := range busNums {
		leaveTime := ((earliestTime / busid) + 1) * busid
		// fmt.Printf("bus: %v leave:%v\n", busid, leaveTime)
		if leaveTime < minLeaveTime {
			minLeaveTime = leaveTime
			minLeaveBus = busid
		}
	}

	waitTime := minLeaveTime - earliestTime
	result = minLeaveBus * waitTime
	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
