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

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p Passport) isValid() bool {
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if p.byr == "" || len(p.byr) < 4 {
		return false
	} else {
		i, err := strconv.ParseInt(p.byr, 10, 64)
		if err != nil {
			return false //Not a number
		}
		if i < 1920 || i > 2002 {
			// fmt.Println("Invalid byr")
			return false
		}
	}
	if p.iyr == "" || len(p.iyr) < 4 {
		//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		return false
	} else {
		i, err := strconv.ParseInt(p.iyr, 10, 64)
		if err != nil {
			return false //Not a number
		}
		if i < 2010 || i > 2020 {
			// fmt.Println("Invalid iyr")
			return false
		}
	}
	if p.eyr == "" || len(p.eyr) < 4 {
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		return false
	} else {
		i, err := strconv.ParseInt(p.eyr, 10, 64)
		if err != nil {
			return false //Not a number
		}
		if i < 2020 || i > 2030 {
			// fmt.Println("Invalid eyr")
			return false
		}
	}
	if p.hgt == "" {
		// hgt (Height) - a number followed by either cm or in:
		// If cm, the number must be at least 150 and at most 193.
		// If in, the number must be at least 59 and at most 76.
		return false
	} else {
		r := regexp.MustCompile("([0-9]+)(in|cm)")
		m := r.FindStringSubmatch(p.hgt)
		// fmt.Println(m)
		if m == nil {
			//didn't match the correct format for the height
			return false
		}

		units := m[2]
		if units != "cm" && units != "in" {
			return false
		}

		i, err := strconv.ParseInt(m[1], 10, 64)
		if err != nil {
			return false //Not a number
		}
		if units == "cm" && (i < 150 || i > 193) {
			return false
		}
		if units == "in" && (i < 59 || i > 76) {
			return false
		}
	}
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if p.hcl == "" {
		return false
	} else {
		r := regexp.MustCompile("^#[0-9a-f]{6}$")

		if !r.MatchString(p.hcl) {
			return false
		}
	}
	if p.ecl == "" {
		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.

		return false
	} else {
		r := regexp.MustCompile("(amb|blu|brn|gry|grn|hzl|oth)")
		if !r.MatchString(p.ecl) {
			return false
		}
	}
	if p.pid == "" {
		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		return false
	} else {
		r := regexp.MustCompile("^([0-9]){9}$")

		if !r.MatchString(p.pid) {
			return false
		}
	}
	// cid (Country ID) - ignored, missing or not.

	return true
}

func (p *Passport) addDetails(passportLine string) {
	segments := strings.Split(passportLine, " ")

	for _, kv := range segments {
		//Should now be in the format 'field:value'
		parts := strings.Split(kv, ":")
		if parts[0] == "byr" {
			p.byr = parts[1]
		}
		if parts[0] == "iyr" {
			p.iyr = parts[1]
		}
		if parts[0] == "eyr" {
			p.eyr = parts[1]
		}
		if parts[0] == "hgt" {
			p.hgt = parts[1]
		}
		if parts[0] == "hcl" {
			p.hcl = parts[1]
		}
		if parts[0] == "ecl" {
			p.ecl = parts[1]
		}
		if parts[0] == "pid" {
			p.pid = parts[1]
		}
		if parts[0] == "cid" {
			p.cid = parts[1]
		}
	}
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

func processInput(f io.Reader) string {
	startTime := time.Now().Unix()
	result := 0

	//Passport to check
	passport := Passport{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		if s.Text() == "" {
			// fmt.Println("End of Passport, check validity")
			if passport.isValid() {
				result++
				// fmt.Println("Valid:", passport, result)
			}
			// fmt.Println("Invalid:", passport, result)
			//Reset to an empty passport...
			passport = Passport{}
		} else {
			passport.addDetails(s.Text())
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}
	//Need to do one last passport check at the end of the input
	if passport.isValid() {
		result++
	}

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result %v in %d seconds\n", result, endTime-startTime)

	return fmt.Sprintf("%d", result)
}
