package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./aoc-d3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var result [12]int64
	var revResult [12]int64
	var linecount (int64)
	resultString := ""
	revResultString := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linecount++
		line := scanner.Text()
		for i := 0; i < 12; i++ {
			if line[i] == 48 {
				result[i]++
			}
		}
	}
	for i, element := range result {

		if element-linecount/2 > 0 {
			result[i] = 1
			revResult[i] = 0
		} else {
			result[i] = 0
			revResult[i] = 1
		}
		revResultString = revResultString + strconv.FormatInt(revResult[i], 2)
		resultString = resultString + strconv.FormatInt(result[i], 2)
	}
	gamma, _ := strconv.ParseUint(resultString, 2, 16)
	epsilon, _ := strconv.ParseUint(revResultString, 2, 16)
	fmt.Printf("\n*************\nGamma: %d\nEpsilon: %d \nResult: %d\n*************\n ", gamma, epsilon, gamma*epsilon)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
