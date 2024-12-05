package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// this is probably slow and you could speed it up by testing smaller and smaller segments to see if each of them are safe

func abs(i int) int{
	if i >= 0{
		return i
	} else {
		return i * -1
	}
}

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

func isSafe(report []string) bool {
	if len(report) < 2 {
		return true
	}

	doesIncrease := toIntCheckErr(report[0]) < toIntCheckErr(report[1])

	for i := 1; i < len(report); i++ {
		reportIMin := toIntCheckErr(report[i-1])
		reportI := toIntCheckErr(report[i])

		if (!(reportIMin < reportI) && doesIncrease) || (reportIMin < reportI && !doesIncrease) {
			return false
		}
		if abs(reportIMin - reportI) > 3 || abs(reportIMin - reportI) < 1 {
			return false
		}
	}

	return true
}

func main() {
	debug := false 

	file, err := os.Open("./puzzleInput.txt")
	defer file.Close()

	check(err)

	s := bufio.NewScanner(file)

	sum := 0

	for s.Scan(){
		report := strings.Split(s.Text(), " ")

		reportSafe := isSafe(report)

		if reportSafe{
			sum += 1
		} 
		if debug {
			fmt.Println(report, reportSafe)
		}

	}

	fmt.Printf("Safe Count: %d\n", sum)
}