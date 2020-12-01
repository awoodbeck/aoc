package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	input := flag.Arg(0)
	if input == "" {
		log.Fatal("input file required")
	}

	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("opening input file: %v", err)
	}
	defer func() {
		if cErr := file.Close(); cErr != nil {
			log.Printf("closing input file: %v", cErr)
		}
	}()

	expenses := make([]int, 0, 2)
	r := bufio.NewReader(file)

OUTER:
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			// We should account for the lack of a trailing new line
			// in the file by only breaking if we received an EOF and
			// the string is empty.
			if err == io.EOF && s == "" {
				log.Print("reached EOF")
				break
			}

			if err != io.EOF {
				log.Fatalf("reading input: %v", err)
			}
		}

		expense, cErr := strconv.Atoi(strings.TrimRight(s, "\n"))
		if cErr != nil {
			log.Printf("atoi: %v", cErr)
			continue
		}

		for _, e := range expenses {
			if e+expense == 2020 {
				log.Printf("%d * %d = %d", e, expense, e*expense)
				break OUTER
			}
		}

		expenses = append(expenses, expense)
	}
}
