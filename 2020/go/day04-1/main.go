package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/awoodbeck/aoc/2020/go/input"
)

type passport struct {
	BirthYear,
	IssueYear,
	ExpirationYear,
	Height,
	HairColor,
	EyeColor,
	PassportID,
	CountryID string
}

func (p *passport) Valid(b string) error {
	// flatten the elements by cleaning up the white space around each one
	flat := strings.Split(strings.ReplaceAll(strings.TrimSpace(b), "\n", " "), " ")

	if len(flat) == 0 {
		return fmt.Errorf("empty passport details")
	}

	for _, f := range flat {
		s := strings.Split(f, ":")
		switch k, v := s[0], s[1]; k {
		case "byr":
			p.BirthYear = v
		case "iyr":
			p.IssueYear = v
		case "eyr":
			p.ExpirationYear = v
		case "hgt":
			p.Height = v
		case "hcl":
			p.HairColor = v
		case "ecl":
			p.EyeColor = v
		case "pid":
			p.PassportID = v
		case "cid":
			p.CountryID = v
		}
	}

	var err error
	switch {
	case p.BirthYear == "":
		err = fmt.Errorf("missing birth year")
	case p.IssueYear == "":
		err = fmt.Errorf("missing issue year")
	case p.ExpirationYear == "":
		err = fmt.Errorf("missing expiration year")
	case p.Height == "":
		err = fmt.Errorf("missing height")
	case p.HairColor == "":
		err = fmt.Errorf("missing hair color")
	case p.EyeColor == "":
		err = fmt.Errorf("missing eye color")
	case p.PassportID == "":
		err = fmt.Errorf("missing passport ID")
	}

	return err
}

func main() {
	data, err := input.ReadAll(4)
	if err != nil {
		log.Fatal(err)
	}

	var validPassports int
	for _, d := range bytes.Split(data, []byte("\n\n")) {
		if (&passport{}).Valid(string(d)) == nil {
			validPassports++
		}
	}

	log.Printf("%d valid passports", validPassports)
}
