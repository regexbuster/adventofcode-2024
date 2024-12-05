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

	for i := 0; i < len(search); i++ {
		// fmt.Println(search[i])
		for j := 0; j < len(search[i]); j++ {
			canHorizSearch := (len(search[i]) - j) >= 4
			canVertSearch := (len(search) - i) >= 4
			canDiagLeftSearch := canVertSearch && (j >= 3)
			canDiagRightSearch := canVertSearch && canHorizSearch

			if canHorizSearch {
				potentialValue := strings.Join(search[i][j:j+4], "")
				if  potentialValue == "XMAS" || potentialValue == "SAMX" {
					fmt.Printf("HORIZ %s at (%d, %d) to (%d, %d)\n", potentialValue, i, j, i, j+3)
					sum += 1
				}
			}

			if canVertSearch {
				potentialValue := strings.Join([]string{search[i][j], search[i+1][j], search[i+2][j], search[i+3][j]}, "")
				if  potentialValue == "XMAS" || potentialValue == "SAMX" {
					fmt.Printf("VERT  %s at (%d, %d) to (%d, %d)\n", potentialValue, i, j, i+3, j)
					sum += 1
				}
			}

			if canDiagLeftSearch {
				potentialValue := strings.Join([]string{search[i][j], search[i+1][j-1], search[i+2][j-2], search[i+3][j-3]}, "")
				if potentialValue == "XMAS" || potentialValue == "SAMX" {
					fmt.Printf("DRIGT  %s at (%d, %d) to (%d, %d)\n", potentialValue, i, j, i+3, j-3)
					sum += 1
				}
			}

			if canDiagRightSearch {
				potentialValue := strings.Join([]string{search[i][j], search[i+1][j+1], search[i+2][j+2], search[i+3][j+3]}, "")
				if potentialValue == "XMAS" || potentialValue == "SAMX" {
					fmt.Printf("DLEFT  %s at (%d, %d) to (%d, %d)\n", potentialValue, i, j, i+3, j+3)
					sum += 1
				}
			}	
		}
	}

	// fmt.Println(search)

	fmt.Println(len(search), len(search[0]))

	fmt.Printf("Sum of XMAS: %d\n", sum)
}