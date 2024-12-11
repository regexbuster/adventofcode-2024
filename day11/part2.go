package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	// "slices"
)

func check(e error) {
	if e != nil {
		check(e)
	}
}

func digitsEven (i int) bool {
	s := strconv.Itoa(i)

	return len(strings.Split(s, "")) % 2 == 0
}

// func (sm *StackManager) blink (n *Node) []Node {
// 	var r []Node

// 	if n.value == 0 {
// 		r = append(r, Node{value:1, depth:(n.depth + 1)})
// 	} else if digitsEven(n.value) {
// 		s := strings.Split(strconv.Itoa(n.value), "")
		
// 		pOne, errOne := strconv.Atoi(strings.Join(s[:len(s) / 2], ""))
// 		check(errOne)

// 		pTwo, errTwo := strconv.Atoi(strings.Join(s[len(s) / 2:], ""))
// 		check(errTwo)

// 		r = append(r, Node{value:pTwo, depth:(n.depth + 1)})
// 		r = append(r, Node{value:pOne, depth:(n.depth + 1)})
// 	} else {	
// 		r = append(r, Node{value:(n.value * 2024), depth:(n.depth + 1)})
// 	}

// 	return r
// }

type Node struct {
	val		int
	blinks	int
}

type RockManager struct {
	active		[]chan []Node
	future		[]Node

	maxActive	int
	maxBlinks	int

	count		int
}

func (rm *RockManager) manage() {
	// while there are rocks to blink at
	for !(len(rm.active) == 0 && len(rm.future) == 0) {
		// if there is room for more rocks to be observed; else
		if len(rm.active) < rm.maxActive && len(rm.future) > 0{
			c := make(chan []Node) 
			rm.active = append(rm.active, c)

			go func(n Node, cc chan []Node){
				nn := []Node{}
				cc <- nn
			}(rm.future[0], c)

			rm.future = rm.future[1:]
			fmt.Printf("%T\n", rm.active)
			fmt.Println(rm.future)
		} else {
			// fmt.Println(rm.active)
			for i := 0; i < len(rm.active); i++ {
				select {
				case ns := <-rm.active[i]:
					fmt.Println(i, "hello")
					rm.future = append(rm.future, ns...)
					rm.active = append(rm.active[:i], rm.active[i+1:]...)
					i--
				default:
					// fmt.Println(i, "bye")
					rm.count++
				}
			}
			// break
		}

		// if rm.count > 10 {
		// 	break
		// }
	}
}

func main() {
	file, err := os.Open("./testInput.txt")
	check(err)

	s := bufio.NewScanner(file)

	s.Scan()

	list := strings.Split(s.Text(), " ")

	// var intList []int

	rm := RockManager{maxActive:2, maxBlinks:25}

	for _, v := range list {
		intV, errV := strconv.Atoi(v)
		check(errV)

		// intList = append(intList, intV)
		rm.future = append(rm.future, Node{val:intV, blinks:0})
	}

	fmt.Println(rm.future)

	rm.manage()

	// channels := make([]chan []int, len(intlist))

	// for j := 0; j < len(intlist); j++ {
	// 	channels[j] = make(chan []int)
	// }

	// for i, v := range intlist {
	// 	go func(v int, c chan []int) {
	// 		l := []int{v}

	// 		for i := 0; i < 25; i++ {
	// 			blink(&l)
	
	// 			fmt.Println(i, v, len(l))
	// 		}

	// 		c <- l
	// 	}(v, channels[i])
	// }

	// var sum int

	// for j := 0; j < len(intlist); j++ {
	// 	sum = sum + len(<-channels[j])
	// }

	// fmt.Println(sum)
}