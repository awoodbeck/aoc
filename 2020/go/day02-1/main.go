package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/awoodbeck/aoc/2020/go/input"
)

var re = regexp.MustCompile(`^([0-9]+)-([0-9]+)\s([a-z]):\s([a-z]+)$`)

func main() {
	data, err := input.Read(2)
	if err != nil {
		log.Fatal(err)
	}

	var (
		valid int
		v     = new(validator)
	)
	for i, d := range data {
		if err := v.Parse(d); err != nil {
			log.Printf("line %d: %v", i+1, err)
			continue
		}

		if v.Valid() {
			valid++
		}
	}

	log.Println("valid passwords:", valid)
}

type validator struct {
	min, max         int
	letter, password string
}

func (v *validator) Parse(p string) error {
	m := re.FindStringSubmatch(p)
	if l := len(m); l < 5 { // match + 4 sub matches
		return fmt.Errorf("invalid password policy string: %q", p)
	}

	var err error
	v.min, err = strconv.Atoi(m[1])
	if err != nil {
		return err
	}

	v.max, err = strconv.Atoi(m[2])
	if err != nil {
		return err
	}

	v.letter = m[3]
	v.password = m[4]

	return nil
}

func (v validator) Valid() bool {
	count := make(map[string]int)

	for _, r := range v.password {
		count[string(r)]++
	}

	c, ok := count[v.letter]

	return ok && c >= v.min && c <= v.max
}
