package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

type ValueList struct {
	Values []int
}

func check(e error) {
	if e != nil {
		check(e)
	}
}

func digitsEven (i int) bool {
	s := strconv.Itoa(i)

	return len(strings.Split(s, "")) % 2 == 0
}

func (l *ValueList) add (i int) {
	l.Values = append(l.Values, i)
}

func (l *ValueList) blink () {
	var r ValueList

	for _, v := range l.Values {
		if v == 0 {
			r.add(1)
		} else if digitsEven(v) {
			s := strings.Split(strconv.Itoa(v), "")
			
			pOne, errOne := strconv.Atoi(strings.Join(s[:len(s) / 2], ""))
			check(errOne)

			pTwo, errTwo := strconv.Atoi(strings.Join(s[len(s) / 2:], ""))
			check(errTwo)

			r.add(pOne)
			r.add(pTwo)
		} else {
			r.add(v * 2024)
		}
	}

	l.Values = slices.Clone(r.Values)
}

func main() {
	file, err := os.Open("./testInput.txt")
	check(err)

	s := bufio.NewScanner(file)

	s.Scan()

	text := s.Text()

	list := strings.Split(text, " ")

	var intlist ValueList

	for _, v := range list {
		intV, errV := strconv.Atoi(v)
		check(errV)

		intlist.Values = append(intlist.Values, intV)
	}

	fmt.Println(intlist.Values)

	for i := 0; i < 25; i++ {
		intlist.blink()

		fmt.Printf("Blink #%d:	%d\n", i + 1, len(intlist.Values))
	}
}