package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	// "slices"
	"strconv"
)

// if err, panic it 
func check(e error){
	if e != nil {
		panic(e)
	}
}

func toIntCheckErr(s string) int{
	val, err := strconv.Atoi(s)
	check(err)

	return val
}

func main() {
	file, err := os.Open("./puzzleInput.txt")
	defer file.Close()

	check(err)

	s := bufio.NewScanner(file)

	rangeRegex, rangeErr := regexp.Compile(`do\(\).*?don't\(\)`)
	check(rangeErr)

	var program string

	var sum int

	for s.Scan() {
		text := s.Text()
		program += text
	}

	program = "do()" + program + "don't()"

	rangeIndexes := rangeRegex.FindAllStringIndex(program, -1)

	multRegex, multErr := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	check(multErr)

	for _, val := range rangeIndexes {
		mults := multRegex.FindAllStringSubmatch(program[val[0]:val[1]], -1)

		for _, multVal := range mults {
			sum = sum + toIntCheckErr(multVal[1]) * toIntCheckErr(multVal[2])
		}
	}

	fmt.Printf("Mul Sum: %d\n", sum)
}