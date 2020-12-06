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
		affirmap := make(map[string]struct{})
		for _, answer := range strings.Split(strings.ReplaceAll(group, "\n", ""), "") {
			affirmap[answer] = struct{}{}
		}

		affirmatives += len(affirmap)
	}

	log.Printf("total affirmative answers = %d", affirmatives)
}
