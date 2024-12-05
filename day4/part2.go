package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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

	var sum int
	var search [][]string

	for s.Scan(){
		search = append(search, strings.Split(s.Text(), ""))
	}

	for i := 0; i < len(search) - 2; i++ {
		for j := 0; j < len(search[i]) - 2; j++ {
			testBlock := [][]string{search[i][j:j+3], search[i+1][j:j+3], search[i+2][j:j+3]}

			// if the potential block doesn't have an A in the center it wont work so end early
			if testBlock[1][1] != "A" {
				continue
			}

			if testBlock[0][0] == "M" && testBlock[0][2] == "M" && testBlock[2][0] == "S" && testBlock[2][2] == "S" {
				fmt.Printf("UBLOCK %s at (%d, %d) to (%d, %d)\n", testBlock, i, j, i+2, j+2)
					sum += 1
			}

			if testBlock[0][0] == "M" && testBlock[2][0] == "M" && testBlock[0][2] == "S" && testBlock[2][2] == "S" {
				fmt.Printf("LBLOCK %s at (%d, %d) to (%d, %d)\n", testBlock, i, j, i+2, j+2)
					sum += 1
			}

			if testBlock[2][0] == "M" && testBlock[2][2] == "M" && testBlock[0][0] == "S" && testBlock[0][2] == "S" {
				fmt.Printf("DBLOCK %s at (%d, %d) to (%d, %d)\n", testBlock, i, j, i+2, j+2)
					sum += 1
			}

			if testBlock[0][2] == "M" && testBlock[2][2] == "M" && testBlock[0][0] == "S" && testBlock[2][0] == "S" {
				fmt.Printf("RBLOCK %s at (%d, %d) to (%d, %d)\n", testBlock, i, j, i+2, j+2)
					sum += 1
			}
	
			// fmt.Println(testBlock)
			
		}

	}

	// fmt.Println(search)

	fmt.Println(len(search), len(search[0]))

	fmt.Printf("Sum of XMAS: %d\n", sum)
}