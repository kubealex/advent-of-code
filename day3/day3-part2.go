package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findOxygen(list []string, index int) string {
	bytecode := calculateRelevant(list, index, "oxy")
	var updatedList []string
	for _, element := range list {
		if element[index] == bytecode {
			updatedList = append(updatedList, element)
			fmt.Println(updatedList)
		}
	}
	if index < 11 && len(updatedList) > 1 {
		index++
		findOxygen(updatedList, index)
	}
	return updatedList[0]
}
func findCO2(list []string, index int) string {
	bytecode := calculateRelevant(list, index, "co2")
	var updatedList []string
	fmt.Println(updatedList)

	for _, element := range list {
		if element[index] == bytecode {
			updatedList = append(updatedList, element)
			fmt.Println(updatedList)
		}
	}
	if index < 11 && len(updatedList) > 1 {
		index++
		findCO2(updatedList, index)
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
	fmt.Println(one)
	fmt.Println(zero)

	if (kind == "oxy" && (one >= zero)) || (kind == "co2" && zero > one) {
		return 49
	} else if (kind == "oxy" && zero > one) || (kind == "co2" && (one >= zero)) {
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
	oxygen := findOxygen(codeList, 0)
	co2 := findCO2(codeList, 0)
	oxyfinal, _ := strconv.ParseInt(oxygen, 2, 32)
	co2final, _ := strconv.ParseInt(co2, 2, 32)
	println(oxyfinal * co2final)
	/* 	for i, element := range result {

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
	   	fmt.Printf("Result: \n Gamma: %d\n Epsilon: %d \n Result: %d\n ", gamma, epsilon, gamma*epsilon)
	   	if err := scanner.Err(); err != nil {
	   		log.Fatal(err)
	   	} */
}
