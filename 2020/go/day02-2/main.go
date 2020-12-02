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
	pos1, pos2       int
	letter, password string
}

func (v *validator) Parse(p string) error {
	m := re.FindStringSubmatch(p)
	if l := len(m); l < 5 { // match + 4 sub matches
		return fmt.Errorf("invalid password policy string: %q", p)
	}

	var err error
	v.pos1, err = strconv.Atoi(m[1])
	if err != nil {
		return err
	}

	v.pos2, err = strconv.Atoi(m[2])
	if err != nil {
		return err
	}

	v.letter = m[3]
	v.password = m[4]

	return nil
}

func (v validator) Valid() bool {
	p1 := string(v.password[v.pos1-1]) == v.letter
	p2 := string(v.password[v.pos2-1]) == v.letter

	return p1 != p2
}
