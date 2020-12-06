package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countAnswers(group []string) (int, int) {
	answers := map[rune]int{}
	for _, a := range group {
		for _, r := range a {
			answers[r]++
		}
	}

	allYes := 0
	for _, a := range answers {
		if a == len(group) {
			allYes++
		}
	}
	return len(answers), allYes
}

func day6() {
	file, err := os.Open("./../inputs/6.txt")
	if err != nil {
		fmt.Printf("couldn't open file: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	groups := [][]string{}
	group := []string{}

	for scanner.Scan() {
		if line := scanner.Text(); strings.TrimSpace(line) != "" {
			group = append(group, line)
		} else {
			groups = append(groups, group)
			group = []string{}
		}
	}
	groups = append(groups, group)

	sum, sumAllYes := 0, 0
	for _, g := range groups {
		s, sAll := countAnswers(g)
		sum += s
		sumAllYes += sAll
	}

	fmt.Printf("sum of answers across groups: %d, %d\n", sum, sumAllYes)
}
