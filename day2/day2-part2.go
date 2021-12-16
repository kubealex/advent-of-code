package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./aoc-d2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	coords := [3]int{0, 0, 0} // position, depth, aim
	scanner := bufio.NewScanner(file)
	re_command := regexp.MustCompile("down|forward|up")
	re_position := regexp.MustCompile("(\\d)+")
	for scanner.Scan() {
		command := re_command.FindString(scanner.Text())
		position, _ := strconv.Atoi(re_position.FindString(scanner.Text()))
		switch command {
		case "down":
			coords[2] += position
		case "up":
			coords[2] -= position
		case "forward":
			coords[0] += position
			coords[1] = coords[1] + (coords[2] * position)
		}
	}

	fmt.Printf("\n*************\nResult: %d\n*************\n ", coords[0]*coords[1])
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
