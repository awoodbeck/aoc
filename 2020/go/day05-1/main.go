/*
I'd like to point out that solving this puzzle is as simple as converting
B/R to 1 and F/L to 0 then treating the result as a binary number.

Taking their first example of BFFFBBFRRR equaling seat ID 567:
	BFFFBBF == 1000110 == 70
	RRR == 111 == 7
	Seat ID == 70*8+7 == 567

But I don't often get to write binary searches, so I'm going to implement the
solution the long way.
*/

package main

import (
	"log"

	"github.com/awoodbeck/aoc/2020/go/input"
)

var (
	highestSeatID int
	rows          = make([]int, 128)
	seats         = make([]int, 8)
)

func main() {
	data, err := input.ReadBytes(5)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range data {
		row := findInt(rows, b[:7])
		seat := findInt(seats, b[7:])

		if id := row*8 + seat; id > highestSeatID {
			highestSeatID = id
		}
	}

	log.Printf("highest ID = %d", highestSeatID)
}

func findInt(a []int, b []byte) int {
	if l := len(a); l > 1 {
		mid := l / 2

		switch dir := b[0]; dir {
		case 'F', 'L': // left
			return findInt(a[:mid], b[1:])
		case 'B', 'R': // right
			return findInt(a[mid:], b[1:])
		default:
			panic("ran out of directions")
		}
	}

	return a[0]
}

func init() {
	for i := 0; i < 128; i++ {
		rows[i] = i
	}

	for i := 0; i < 8; i++ {
		seats[i] = i
	}
}
