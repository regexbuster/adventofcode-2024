package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"slices"
	"strconv"
)

// it think there's a abs function in a standard import but I don't want the whole package
func abs(i int) int{
	if i >= 0{
		return i
	}
	
	return i * -1
}

// if err, panic it 
func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./testInput.txt")
	defer file.Close()

	check(err)

	s := bufio.NewScanner(file)

	// holds the values for each list
	var listOne []int
	var listTwo []int

	// scan through the lists and adds each side to their above slice
	for s.Scan() {
		splitted := strings.Split(s.Text(), "   ")

		g, errg := strconv.Atoi(splitted[0])
		check(errg)
		listOne = append(listOne, g)

		h, errh := strconv.Atoi(splitted[1])
		check(errh)
		listTwo = append(listTwo, h)
	}

	// sort since we need smallest to largest
	slices.Sort(listOne)
	slices.Sort(listTwo)

	sum := 0

	// convert strings to int and then add the difference abs(num1 - num2)
	for i := 0; i < len(listOne); i++ {
		sum += (abs(listOne[i]-listTwo[i]))
	}

	fmt.Printf("Total Distance: %d\n", sum)
}