package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/awoodbeck/aoc/2020/go/input"
)

var (
	reHeight     = regexp.MustCompile(`^([0-9]+)(cm|in)$`)
	reHairColor  = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	reEyeColor   = regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	rePassportID = regexp.MustCompile(`^[0-9]{9}$`)
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

	if by, err := strconv.Atoi(p.BirthYear); err != nil ||
		len(p.BirthYear) != 4 || by < 1920 || by > 2002 {
		return fmt.Errorf("invalid birth year")
	}

	if iy, err := strconv.Atoi(p.IssueYear); err != nil ||
		len(p.IssueYear) != 4 || iy < 2010 || iy > 2020 {
		return fmt.Errorf("invalid issue year")
	}

	if ey, err := strconv.Atoi(p.ExpirationYear); err != nil ||
		len(p.ExpirationYear) != 4 || ey < 2020 || ey > 2030 {
		return fmt.Errorf("invalid issue year")
	}

	if ey, err := strconv.Atoi(p.ExpirationYear); err != nil ||
		len(p.ExpirationYear) != 4 || ey < 2020 || ey > 2030 {
		return fmt.Errorf("invalid issue year")
	}

	h := reHeight.FindStringSubmatch(p.Height)
	if h == nil || len(h) != 3 {
		return fmt.Errorf("invalid height")
	} else if hn, err := strconv.Atoi(h[1]); err != nil ||
		(h[2] == "cm" && (hn < 150 || hn > 193)) ||
		(h[2] == "in" && (hn < 59 || hn > 76)) {
		return fmt.Errorf("invalid height")
	}

	if !reHairColor.MatchString(p.HairColor) {
		return fmt.Errorf("invalid hair color")
	}

	if !reEyeColor.MatchString(p.EyeColor) {
		return fmt.Errorf("invalid eye color")
	}

	if !rePassportID.MatchString(p.PassportID) {
		return fmt.Errorf("invalid passport ID")
	}

	return nil
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
