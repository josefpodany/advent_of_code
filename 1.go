package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum2(arr []int) (int, int) {
	for _, fst := range arr {
		for _, snd := range arr {
			if fst+snd == 2020 {
				return fst, snd
			}
		}
	}
	return 0, 0
}

func sum3(arr []int) (int, int, int) {
	for _, fst := range arr {
		for _, snd := range arr {
			for _, thd := range arr {
				if fst+snd+thd == 2020 {
					return fst, snd, thd
				}
			}
		}
	}
	return 0, 0, 0
}

func main() {
	// load the inputs
	file, err := os.Open("./inputs/1.txt")
	if err != nil {
		fmt.Printf("couldn't open file: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	expenses := []int{}
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("couldn't parse input: %s\n", err.Error())
		}
		expenses = append(expenses, i)
	}

	fst, snd := sum2(expenses)
	fmt.Printf("%d * %d = %d\n", fst, snd, fst*snd)
	var thd int
	fst, snd, thd = sum3(expenses)
	fmt.Printf("%d * %d * %d = %d\n", fst, snd, thd, fst*snd*thd)
}
