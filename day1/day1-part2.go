package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkIncrement(prev int, curr int, incr *int) {
	if curr > prev {
		*incr++
	}
}

func main() {
	file, err := os.Open("./aoc-d1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	increment := new(int)
	prevSum := 0
	window := [3]int{0, 0, 0}
	scanner := bufio.NewScanner(file)
	for i := 0; i <= 2; i++ {
		scanner.Scan()
		window[i], _ = strconv.Atoi(scanner.Text())
		prevSum += window[i]
	}

	for scanner.Scan() {
		window[0] = window[1]
		window[1] = window[2]
		window[2], _ = strconv.Atoi(scanner.Text())
		sum := window[0] + window[1] + window[2]
		checkIncrement(prevSum, sum, increment)
		prevSum = sum
	}

	fmt.Printf("\n*************\nTotal increments: %d\n*************\n ", *increment)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
