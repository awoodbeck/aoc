package main

import (
	"log"

	"github.com/awoodbeck/aoc/2020/go/input"
)

var (
	seenIDs = make(map[int]struct{})
	rows    = make([]int, 128)
	seats   = make([]int, 8)
)

func main() {
	data, err := input.ReadBytes(5)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range data {
		row := findInt(rows, b[:7])
		seat := findInt(seats, b[7:])
		seenIDs[row*8+seat] = struct{}{}
	}

	var seatID int

	// I start searching for the missing seat ID in the middle
	// because I don't know which seats and the front and back
	// of the plane do not exist.
	//
	// Also, I don't want to assume that my seat ID is less than
	// the highest seat ID I saw in part 1. Therefore, I'll search
	// for my seat across the entire range of possible seat IDs.
	// This should ensure I find my seat even if it happens to be
	// the first or last one on the plane while considering the
	// condition that some of the seats at the head and tail of the
	// plane may not exist.
	mid := (127*8 + 7) / 2
	for i := 0; i < mid; i++ {
		if _, ok := seenIDs[mid+i]; !ok {
			seatID = mid + i
			break
		}
		if _, ok := seenIDs[mid-i]; !ok {
			seatID = mid - i
			break
		}
	}

	log.Printf("my seat ID = %d", seatID)
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
