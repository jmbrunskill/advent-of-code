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

type sleeper struct {
	id          int
	sleepTime   int
	sleepMinute [60]int
}

func (s sleeper) String() string {
	return fmt.Sprintf("ID: %d, Asleep for %d minutes - %v", s.id, s.sleepTime, s.sleepMinute)
}

func (s *sleeper) addSleep(startMinute, endMinute int) int {
	// fmt.Println("addingSleep - ", s.id, startMinute, endMinute)
	s.sleepTime += (endMinute - startMinute)
	for i := startMinute; i < 60 && i < endMinute; i++ {
		s.sleepMinute[i]++
	}
	// fmt.Println(s)
	return s.sleepTime
}

func (s sleeper) sleepiestMinute() (int, int) {
	maxMin := 0
	maxSleeps := 0
	for i := 0; i < len(s.sleepMinute); i++ {
		if s.sleepMinute[i] > maxSleeps {
			maxSleeps = s.sleepMinute[i]
			maxMin = i
		}

	}
	return maxMin, maxSleeps
}

func processInput(f io.Reader) string {
	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		lines = append(lines, s.Text())

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	sort.Strings(lines)

	//String sort should give us the correct order, now to process them...

	//EXAMPLE INPUT
	//[1518-11-01 00:00] Guard #10 begins shift
	//[1518-11-01 00:05] falls asleep
	//[1518-11-01 00:25] wakes up

	//Keep track of our biggest sleeper
	sleepers := make(map[int]*sleeper)

	var err error
	currentGuard := 0
	currentGuardAwake := true
	lastEventMin := 0

	for _, line := range lines {
		// fmt.Println(line)
		//date := line[1:11]
		hour := line[12:14]
		minute := line[15:17]
		action := line[18:]
		// fmt.Printf("[%v %v:%v] - %v\n", date, hour, minute, action)

		fields := strings.Fields(action)
		switch fields[0] {
		case "Guard":
			id, err := strconv.Atoi(fields[1][1:])
			if err != nil {
				log.Fatalf("could not parse id %q: %v", fields[1][1:], err)
			}

			if sleepers[id] == nil {
				sleepers[id] = &sleeper{id: id}
			}

			if currentGuard != 0 && !currentGuardAwake {
				//last guard must have slept through till midnight
				_ = sleepers[currentGuard].addSleep(lastEventMin, 60)
			}

			currentGuard = id
			currentGuardAwake = true //Start shift awake
			lastEventMin, err = strconv.Atoi(minute)
			if err != nil {
				log.Fatalf("could not parse minute %q: %v", minute, err)
			}
			if hour != "00" {
				//assume this must be before midnight and they started awake
				lastEventMin = 0
			}
		case "falls":
			if !currentGuardAwake {
				log.Fatalf("Guard %d went to sleep up but he's already asleep!", currentGuard)
			}
			//Guard falls asleep, assume they must have been awake
			currentGuardAwake = false
			lastEventMin, err = strconv.Atoi(minute)
			if err != nil {
				log.Fatalf("could not parse minute %q: %v", minute, err)
			}
		case "wakes":
			if currentGuardAwake {
				log.Fatalf("Guard %d woke up but he's not asleep!", currentGuard)
			}

			thisMinute, err := strconv.Atoi(minute)
			if err != nil {
				log.Fatalf("could not parse minute %q: %v", minute, err)
			}
			if currentGuard != 0 && !currentGuardAwake {
				//last guard slept through till midnight
				_ = sleepers[currentGuard].addSleep(lastEventMin, thisMinute)
			}
			lastEventMin = thisMinute
			currentGuardAwake = true

		}
		// fmt.Printf("Guard %v - %v @ %v \n", currentGuard, fields[0], lastEventMin)
	}

	maxSleepMinute := 0
	maxSleepCount := 0
	maxSleepMinuteGuard := 0

	for k := range sleepers {
		// fmt.Println(k, v)
		sleepiestMinute, sleepiestMinuteCount := sleepers[k].sleepiestMinute()
		// fmt.Println(sleepiestMinute, sleepiestMinuteCount)
		if sleepiestMinuteCount > maxSleepCount {
			maxSleepCount = sleepiestMinuteCount
			maxSleepMinuteGuard = k
			maxSleepMinute = sleepiestMinute
		}
	}

	fmt.Printf("Sleepiest Sleeper Minute is ... %d @ %d with %d sleeps\n", maxSleepMinuteGuard, maxSleepMinute, maxSleepCount)

	return fmt.Sprintf("%v", maxSleepMinuteGuard*maxSleepMinute)
}
