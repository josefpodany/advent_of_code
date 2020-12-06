package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func fromBinary(code string, lower rune) int {
	lo, hi := float64(0), math.Pow(2, float64(len(code)))-1
	for _, ch := range code[:len(code)-1] {
		if ch == lower {
			hi = math.Floor((lo + hi) / 2)
		} else {
			lo = math.Ceil((lo + hi) / 2)
		}

	}
	if rune(code[len(code)-1]) == lower {
		return int(lo)
	}
	return int(hi)
}

func day5() {
	file, err := os.Open("./../inputs/5.txt")
	if err != nil {
		fmt.Printf("couldn't open file: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	highest := int(math.Inf(-1))
	allIDs := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		row, column := line[:len(line)-3], line[len(line)-3:]
		ID := fromBinary(row, 'F')*8 + fromBinary(column, 'L')
		allIDs = append(allIDs, ID)
		if ID > highest {
			highest = ID
		}
	}
	fmt.Printf("highest ID: %d\n", highest)

	// We can "sort" the IDs in O(n), since IDs are unique,
	// meaning we can just iterate over the IDs
	// and use the ID as an index to the sorted slice
	sortedIDs := make([]bool, highest+1)
	for _, id := range allIDs {
		sortedIDs[id] = true

	}

	for i := 1; i < len(sortedIDs)-2; i++ {
		if !sortedIDs[i] && sortedIDs[i-1] && sortedIDs[i+1] {
			fmt.Printf("I should be seated at %d\n", i)
		}
	}
}
