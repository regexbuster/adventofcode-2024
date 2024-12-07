package main

import (
	"os"
	"bufio"
	"regexp"
	"fmt"
	"strconv"
)

type OperandType int

const (
	ADD	OperandType = iota
	MUL
)

var operandSymbol = map[OperandType]string{
	ADD: "+",
	MUL: "*",
}

type Equation struct {
	Result int
	Values []int
	Operands []OperandType
}

func NewEquation(result int, values []int) Equation{
	op := make([]OperandType, len(values) - 1)

	for i, _ := range op {
		op[i] = ADD
	}

	return Equation{result, values, op}
}

func (eq Equation) Print() {
	mathEq := strconv.Itoa(eq.Result) + " = " + strconv.Itoa(eq.Values[0])

	for i := 0; i < len(eq.Operands); i++ {
		switch eq.Operands[i]{
		case ADD:
			mathEq = mathEq + operandSymbol[ADD] + strconv.Itoa(eq.Values[i+1])
		case MUL:
			mathEq = mathEq + operandSymbol[MUL] + strconv.Itoa(eq.Values[i+1])
		}
	}

	fmt.Println(mathEq)
}

func (eq Equation) Update() int {
	index := 0
	for index < len(eq.Operands) {
		if eq.Operands[index] == MUL {
			eq.Operands[index] = ADD
			index = index + 1
		} else {
			eq.Operands[index] = MUL
			return index
		}
	}

	return index
}

func (eq Equation) isSolvable() bool {
	for true {
		total := eq.Values[0]
		for i := 0; i < len(eq.Operands); i++ {
			switch eq.Operands[i]{
			case ADD:
				total = total + eq.Values[i+1]
			case MUL:
				total = total * eq.Values[i+1]
			}
		}

		if eq.Result == total {
			return true
		} 

		endIndex := eq.Update()

		if endIndex >= len(eq.Operands){
			return false
		}
	}

	return false
}

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main(){
	file, err := os.Open("./puzzleInput.txt")
	check(err)

	s := bufio.NewScanner(file)

	digitRegex, digitErr := regexp.Compile(`\d+`)
	check(digitErr)

	var sum int
	for s.Scan() {
		text := s.Text()

		values := digitRegex.FindAllString(text, -1)

		var intValues []int
		for _, v := range values {
			intV, intErr := strconv.Atoi(v)
			check(intErr)

			intValues = append(intValues, intV)
		}

		eq := NewEquation(intValues[0], intValues[1:])

		if eq.isSolvable() {
			sum = sum + eq.Result
			eq.Print()
		}
	}

	fmt.Printf("Total Sum: %d\n", sum)
}