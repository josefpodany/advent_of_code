package main

import (
	"bufio"
	"fmt"
	"os"
)

// slideDown with the slope of dx, dy and count how
// many trees will be hit
func slideDown(dx, dy int, levels []string) int {
	width := len(levels[0])
	trees := 0

	for y, x := dy, 0; y < len(levels); y += dy {
		x = (x + dx) % width
		if rune(levels[y][x]) == '#' {
			// oopsie!
			trees++
		}
	}

	return trees
}

// first element represents the dx, second is dy
var slopes = [][]int{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func main() {
	file, err := os.Open("./input.txt")
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

	product := 1
	for _, slope := range slopes {
		dx, dy := slope[0], slope[1]
		trees := slideDown(dx, dy, levels)
		fmt.Printf("slope dx: %d, dy: %d, %d trees will be hit\n", dx, dy, trees)
		product *= trees
	}
	fmt.Printf("product of all trees encountered: %d\n", product)
}
