package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type BingoBoard struct {
}

type BingoCell struct {
}

func Split(r rune) bool {
	return r == ',' || r == ' '
}

func main() {
	file, err := os.Open("./aoc-d4")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var test []string
	for scanner.Scan() {
		line := scanner.Text()
		test = strings.FieldsFunc(line, Split)
		
	}
	fmt.Println(test)
}

//TODO
