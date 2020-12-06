package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	min, max int
	char     rune
	pass     string
}

// hasValidOccurances of a rune according to the sled rental policy
func (l line) hasValidOccurances() bool {
	occurances := 0
	for _, r := range l.pass {
		if r == l.char {
			occurances++
		}
		if occurances > l.max {
			return false
		}
	}
	return occurances >= l.min
}

// hasValidPositions of a rune according to the new job
func (l line) hasValidPositions() bool {
	fst, snd := l.min-1, l.max-1
	return (rune(l.pass[fst]) == l.char) != (rune(l.pass[snd]) == l.char)
}

// countValid passwords based on policies 1 and 2.
func countValid(lines []line) (int, int) {
	occ, pos := 0, 0
	for _, l := range lines {
		if l.hasValidOccurances() {
			occ++
		}

		if l.hasValidPositions() {
			pos++
		}
	}
	return occ, pos
}

func parseBounds(bounds string) (int, int, error) {
	items := strings.Split(bounds, "-")
	if len(items) != 2 {
		return 0, 0, fmt.Errorf("'%s' are not valid bounds", bounds)
	}

	min, err := strconv.Atoi(items[0])
	if err != nil {
		return 0, 0, err
	}

	max, err := strconv.Atoi(items[1])
	if err != nil {
		return 0, 0, err
	}

	return min, max, nil
}

func day2() {
	// load the inputs
	file, err := os.Open("./../inputs/2.txt")
	if err != nil {
		fmt.Printf("couldn't open file: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := []line{}
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		if len(items) != 3 {
			fmt.Printf("err: line '%v' is malformed\n", items)
			continue
		}
		min, max, err := parseBounds(items[0])
		if err != nil {
			fmt.Printf("couldn't parse bounds: %s", err.Error())
			continue
		}

		// get rid of the trailing colon
		if len(items[1]) != 2 {
			fmt.Printf("malformed rune: %s", items[1])
			continue
		}

		lines = append(lines, line{
			min:  min,
			max:  max,
			char: rune(items[1][0]),
			pass: items[2],
		})
	}

	occ, pos := countValid(lines)
	fmt.Printf("valid passwords based on no. of occurances	: %d\n", occ)
	fmt.Printf("valid passwords based on positions: %d\n", pos)

}
