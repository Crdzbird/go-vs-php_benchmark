package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tuple struct {
	X     string
	Y     uint64
	Value string
}

func LoadX(route string) []string {
	file, err := os.Open(route)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func LoadCorpus(route string) ([]Tuple, error) {
	file, err := os.Open(route)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tuples := make([]Tuple, 0)
	for scanner.Scan() {
		entry := strings.Fields(scanner.Text())
		x := entry[0]
		y64, err := strconv.ParseInt(entry[1], 10, 64)
		if err != nil && y64 > 0 {
			return nil, fmt.Errorf("Parsing/conversion error in %x", entry)
		}
		y := uint64(y64)
		value := entry[2]
		tuple := Tuple{X: x, Y: y, Value: value}
		tuples = append(tuples, tuple)
	}

	return tuples, nil
}
