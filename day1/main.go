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
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	current := 0
	previous, _ := strconv.Atoi(scanner.Text())
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		current, _ = strconv.Atoi(scanner.Text())
		checkIncrement(previous, current, increment)
		previous = current
	}
	fmt.Printf("Total increments: %d", *increment)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
