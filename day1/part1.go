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
	} else {
		return i * -1
	}
}

// if err, panic it 
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

	// holds the values for each list
	var listOne []string
	var listTwo []string

	// scan through the lists and adds each side to their above slice
	for s.Scan() {
		splitted := strings.Split(s.Text(), "   ")

		if splitted[0] != ""{
			listOne = append(listOne, splitted[0])
		}

		if splitted[1] != ""{
			listTwo = append(listTwo, splitted[1])
		}
		
	}

	// sort since we need smallest to largest
	slices.Sort(listOne)
	slices.Sort(listTwo)

	sum := 0

	// convert strings to int and then add the difference abs(num1 - num2)
	for i := 0; i < 1000; i++ {
		g, errg := strconv.Atoi(listOne[i])
		check(errg)

		h, errh := strconv.Atoi(listTwo[i])
		check(errh)

		sum += (abs(g-h))
	}

	fmt.Printf("Total Distance: %d\n", sum)
}