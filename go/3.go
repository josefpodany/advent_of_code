package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type result struct {
	dx, dy int
	trees  int
}

// slideDown with the slope of dx, dy and count how
// many trees will be hit
func slideDown(wg *sync.WaitGroup, c chan<- result, dx, dy int, levels []string) {
	defer wg.Done()

	width := len(levels[0])
	trees := 0

	for y, x := dy, 0; y < len(levels); y += dy {
		x = (x + dx) % width
		if rune(levels[y][x]) == '#' {
			trees++
		}
	}

	c <- result{dx: dx, dy: dy, trees: trees}
}

// first element represents the dx, second is dy
var slopes = [][]int{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func consume(c <-chan result) {
	product := 1
	for r := range c {
		fmt.Printf("slope dx: %d, dy: %d, %d trees will be hit\n", r.dx, r.dy, r.trees)
		product *= r.trees
	}
	fmt.Printf("product of all trees encountered: %d\n", product)
}

func day3() {
	file, err := os.Open("./../inputs/3.txt")
	if err != nil {
		fmt.Printf("couldn't open file: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	levels := []string{}
	for scanner.Scan() {
		levels = append(levels, scanner.Text())
	}

	var wg sync.WaitGroup
	c := make(chan result)
	defer close(c)

	go consume(c)
	for _, slope := range slopes {
		wg.Add(1)
		dx, dy := slope[0], slope[1]
		go slideDown(&wg, c, dx, dy, levels)
	}

	wg.Wait()
}
