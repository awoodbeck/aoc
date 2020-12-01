package main

import (
	"bufio"
	"flag"
	"fmt"
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

	expenses, err := getExpenses(file)
	if err != nil {
		log.Fatal(err)
	}

OUTER:
	for _, i := range expenses {
		for _, j := range expenses {
			for _, k := range expenses {
				if i+j+k == 2020 {
					log.Printf("%d * %d * %d = %d", i, j, k, i*j*k)
					break OUTER
				}
			}
		}
	}
}

func getExpenses(r io.Reader) ([]int, error) {
	expenses := make([]int, 0, 3)
	buf := bufio.NewReader(r)

	for {
		s, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF && s == "" {
				break
			}
			if err != io.EOF {
				return nil, fmt.Errorf("reading input: %w", err)
			}
		}

		e, err := strconv.Atoi(strings.TrimRight(s, "\n"))
		if err != nil {
			log.Printf("atoi: %v", err)
			continue
		}

		expenses = append(expenses, e)
	}

	return expenses, nil
}
