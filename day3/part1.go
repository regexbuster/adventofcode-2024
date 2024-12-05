package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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

	mulRegex, regErr := regexp.Compile(`(mul\((\d{1,3}),(\d{1,3})\))`)
	check(regErr)

	var muls []string

	for s.Scan() {
		text := s.Text()
		muls = slices.Concat(muls, mulRegex.FindAllString(text, -1))
	}

	var sum int

	for _, val := range muls {
		numRegex, numErr := regexp.Compile(`(\d{1,3})`)
		check(numErr)

		nums := numRegex.FindAllString(val, -1)

		sum = sum + (toIntCheckErr(nums[0]) * toIntCheckErr(nums[1]))
	}

	fmt.Printf("Mul Sum: %d\n", sum)
}