package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	if p.byr == "" {
		return false
	}
	if p.iyr == "" {
		return false
	}
	if p.eyr == "" {
		return false
	}
	if p.hgt == "" {
		return false
	}
	if p.hcl == "" {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if p.pid == "" {
		return false
	}

	//Skip checking cid *country id
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
			}
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
