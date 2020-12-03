package main

import (
	"log"

	"github.com/awoodbeck/aoc/2020/go/input"
)

const tree = '#'

func main() {
	data, err := input.ReadBytes(3)
	if err != nil {
		log.Fatal(err)
	}

	var trees int

	x := 0
	for y := 0; y < len(data); y++ {
		if data[y][x%len(data[y])] == tree {
			trees++
		}
		x += 3
	}

	log.Printf("🌲 = %d", trees)
}
