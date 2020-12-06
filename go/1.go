package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// sum2 expects the slice to be sorted.
// This way, we can find the sum in O(n) with the
// "two-finger" method.
func sum2(arr []int, sum int) (int, int, bool) {
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left]+arr[right] == sum {
			return arr[left], arr[right], true
		}
		if arr[left]+arr[right] < sum {
			left++
		} else {
			right--
		}
	}
	return 0, 0, false
}

// sum3 expects the slice to be sorted.
// We can find the sum of 3 variables in O(n^2).
func sum3(arr []int, sum int) (int, int, int, bool) {
	for _, thd := range arr {
		if fst, snd, ok := sum2(arr, sum-thd); ok {
			return fst, snd, thd, true
		}
	}
	return 0, 0, 0, false
}

func day1() {
	// load the inputs
	file, err := os.Open("./../inputs/1.txt")
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

	sort.Ints(expenses) // this is O(n*log(n))
	if fst, snd, ok := sum2(expenses, 2020); ok {
		fmt.Printf("%d * %d = %d\n", fst, snd, fst*snd)
	} else {
		fmt.Println("there are no 2 variables that add to 2020")
	}

	if fst, snd, thd, ok := sum3(expenses, 2020); ok {
		fmt.Printf("%d * %d *%d = %d\n", fst, snd, thd, fst*snd*thd)
	} else {
		fmt.Println("there are no 3 variables that add to 2020")
	}
}
