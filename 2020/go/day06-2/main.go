package main

import (
	"log"
	"strings"

	"github.com/awoodbeck/aoc/2020/go/input"
)

func main() {
	data, err := input.ReadAll(6)
	if err != nil {
		log.Fatal(err)
	}

	var affirmatives int
	for _, group := range strings.Split(string(data), "\n\n") {
		affirmap := make(map[string]int)
		for _, answer := range strings.Split(strings.ReplaceAll(group, "\n", ""), "") {
			affirmap[answer]++
		}

		groupSize := len(strings.Split(group, "\n"))
		for _, v := range affirmap {
			if v == groupSize {
				affirmatives++
			}
		}
	}

	log.Printf("total consensus of affirmative answers = %d", affirmatives)
}
