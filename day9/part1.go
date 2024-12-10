package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func repeatedSlice[T comparable](val T, count int) []T {
	var result []T

	for i := 0; i < count; i++ {
		result = append(result, val)
	}

	return result
}

func main() {
	file, err := os.Open("./puzzleInput.txt")
	check(err)

	s := bufio.NewScanner(file)

	var spaces bool
	var counter int
	var unpacked []string

	for s.Scan() {
		diskmap := s.Text()

		for _, v := range strings.Split(diskmap, "") {
			intV, intErr := strconv.Atoi(v)
			check(intErr)

			if spaces {
				unpacked = append(unpacked, repeatedSlice(" ", intV)...)
			} else {
				unpacked = append(unpacked, repeatedSlice(strconv.Itoa(counter), intV)...)
				counter = counter + 1
			}

			spaces = !spaces
		}
	}

	for i := 0; i < len(unpacked); i++ {
		if unpacked[i] != " "{
			continue
		}
	
		for j := len(unpacked) - 1; j > 0; j-- {
			if unpacked[j] == " "{
				continue
			}

			unpacked[i] = unpacked[j]
			unpacked = unpacked[:j]
			break
		}
		// fmt.Println(unpacked)
	}

	var sum int
	for i, v := range unpacked {
		vInt, vErr := strconv.Atoi(v)
		check(vErr)
		sum = sum + (i * vInt)
	}

	fmt.Printf("Checksum: %d\n",sum)
}