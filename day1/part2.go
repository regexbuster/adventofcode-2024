package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./puzzleInput.txt")
	defer file.Close()

	check(err)

	s := bufio.NewScanner(file)

	// hold the first list's values in slice
	// second list is a map so we can have 0 for all values not included in list and probably faster reference than slice (more exact too)
	var listOne []string
	valueMap := make(map[string]int)

	// scan lists added to slice or to map
	for s.Scan() {
		splitted := strings.Split(s.Text(), "   ")

		if splitted[0] != ""{
			listOne = append(listOne, splitted[0])
		}

		if splitted[1] != ""{
			valueMap[splitted[1]] = valueMap[splitted[1]] + 1
		}
	}

	sum := 0

	// convert first list number to int and then multiply it by it's occurance in the second list
	for i := 0; i < 1000; i++ {
		g, errg := strconv.Atoi(listOne[i])
		check(errg)

		sum += (g * valueMap[listOne[i]])
	}

	fmt.Printf("Similarity Score: %d\n", sum)
}