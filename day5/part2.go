package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"slices"
	"strconv"
)

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

	protocols := make(map[string][]string)
	var updates [][]string

	readingProtocols := true

	for s.Scan(){
		text := s.Text()

		if text == ""{
			readingProtocols = false
			continue
		}

		if readingProtocols {
			splitProtocol := strings.Split(text, "|")
			protocols[splitProtocol[1]] = append(protocols[splitProtocol[1]], splitProtocol[0])
		} else {
			updates = append(updates, strings.Split(text, ","))
		}
	}

	var sum int

	for _, iv := range updates {
		var issue bool

		for j, jv := range iv {
			if j >= len(iv) - 1 {
				// break out if too far in as there is nothing after last value
				continue
			}

			issue = slices.ContainsFunc(iv[j+1 : len(iv)], func (n string) bool {
				return slices.Contains(protocols[jv], n)
			})

			if issue {
				// if there's an issue we need to break out to address
				break
			}
		} 

		if issue {
			slices.SortFunc(iv, func(X, Y string) int {
				if slices.Contains(protocols[Y], X) {
					return -1
				}

				return 1
			})

			pos := len(iv) / 2
			sum += toIntCheckErr(iv[pos])
		}
	}

	fmt.Printf("Part 2 Middle Sums: %d\n", sum)
}