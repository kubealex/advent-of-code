package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findParams(list []string, index int, kind string) string {
	bytecode := calculateRelevant(list, index, kind)
	var updatedList []string
	for _, element := range list {
		if element[index] == bytecode {
			updatedList = append(updatedList, element)
		}
	}
	if index < 11 && len(updatedList) > 1 {
		index++
		return findParams(updatedList, index, kind)
	}
	return updatedList[0]
}

func calculateRelevant(list []string, index int, kind string) byte {
	var one int
	var zero int
	for _, element := range list {
		if element[index] == 49 {
			one++
		} else if element[index] == 48 {
			zero++
		}
	}
	if (kind == "oxy" && (one >= zero)) || (kind == "co2" && zero > one) {
		return 49
	} else if (kind == "co2" && (one >= zero)) || (kind == "oxy" && zero > one) {
		return 48
	}
	return 0
}

func main() {
	file, err := os.Open("./aoc-d3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var linecount (int64)
	var codeList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linecount++
		codeList = append(codeList, scanner.Text())
	}
	oxygen := findParams(codeList, 0, "oxy")
	co2 := findParams(codeList, 0, "co2")
	oxyfinal, _ := strconv.ParseInt(oxygen, 2, 32)
	co2final, _ := strconv.ParseInt(co2, 2, 32)
	fmt.Printf("\n*************\nOxygen generator rating: %v \nCO2 scrubber rating: %v \nLife support rating: %v\n*************\n ", oxyfinal, co2final, oxyfinal*co2final)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
