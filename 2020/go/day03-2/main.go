package main

import (
	"log"

	"github.com/awoodbeck/aoc/2020/go/input"
)

const tree = '#'

var slopes = []struct {
	X, Y int
}{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func main() {
	data, err := input.ReadBytes(3)
	if err != nil {
		log.Fatal(err)
	}

	answer := 1

	for _, slope := range slopes {
		trees := 0
		x := 0
		for y := 0; y < len(data); y += slope.Y {
			if data[y][x%len(data[y])] == tree {
				trees++
			}
			x += slope.X
		}

		log.Printf("[%d, %d] 🌲 = %d", slope.X, slope.Y, trees)
		answer *= trees
	}

	log.Printf("answer = %d", answer)
}
