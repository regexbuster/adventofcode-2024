package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"slices"
)

type Board [][]int

func (b Board) getBounds() (int, int) {
	return len(b[0]), len(b)
}

func (b Board) getPosition(p Point) int {
	return b[p.y][p.x]
}

type Point struct {
	x int
	y int
}

type PointQueue struct {
	queue []Point
}

func (pq *PointQueue) enqueue(p Point) {
	pq.queue = append(pq.queue, p)
}

func (pq *PointQueue) dequeue() Point {
	r := pq.queue[0]
	pq.queue = pq.queue[1:]
	return r
}

func (pq PointQueue) peek() Point {
	return pq.queue[0]
}

func (pq PointQueue) rear() Point {
	return pq.queue[len(pq.queue)-1]
}

func (pq PointQueue) isEmpty() bool {
	return len(pq.queue) == 0
}

func (pq PointQueue) size() int {
	return len(pq.queue)
}

func (pq PointQueue) conains(p Point) bool {
	return slices.Contains(pq.queue, p)
}

func (b Board) getTrailheads() []Point {
	var r []Point

	for i, row := range b {
		for j, elem := range row {
			if elem == 0 {
				r = append(r, Point{x:j, y:i})
			}
		}
	}

	return r
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoiSlice(s []string) []int {
	var r []int

	for _, v := range s {
		if v == "."{
			r = append(r, -1)
		} else {
			intV, errV := strconv.Atoi(v)
			check(errV)
			r = append(r, intV)
		}
		
	}

	return r
}

func main() {
	file, err := os.Open("./puzzleInput.txt")
	check(err)

	s := bufio.NewScanner(file)

	var board Board

	for s.Scan() {
		text := s.Text()

		board = append(board, atoiSlice(strings.Split(text, "")))
	}

	trailheads := board.getTrailheads()

	xMax, yMax := board.getBounds()

	var sum int

	for _, v := range trailheads {
		pq := PointQueue{}
		pq.enqueue(v)

		var trailEnds []Point

		for ; !pq.isEmpty() ; {
			pos := pq.dequeue()
			posVal := board.getPosition(pos)
	
			fmt.Println(pq, pos, posVal, pos.y)
	
			if posVal == 9 && !slices.Contains(trailEnds, pos) {
				sum++
				trailEnds = append(trailEnds, pos)
				fmt.Println("found a 9")
				continue
			}
			
			//up
			if (pos.y - 1) >= 0 {
				if posVal + 1 == board.getPosition(Point{x:pos.x, y:(pos.y - 1)}){
					pq.enqueue(Point{x:pos.x, y:(pos.y - 1)})
					fmt.Println("added something :) up")
				}
			}
			//right
			if pos.x + 1 < xMax {
				if posVal + 1 == board.getPosition(Point{x:(pos.x + 1), y:pos.y}){
					pq.enqueue(Point{x:(pos.x + 1), y:pos.y})
					fmt.Println("added something :) right")
				}
			}
			//down
			if pos.y + 1 < yMax {
				if posVal + 1 == board.getPosition(Point{x:pos.x, y:(pos.y + 1)}){
					pq.enqueue(Point{x:pos.x, y:(pos.y + 1)})
					fmt.Println("added something :) down")
				}
			}
			//left
			if pos.x - 1 >= 0 {
				if posVal + 1 == board.getPosition(Point{x:(pos.x - 1), y:pos.y}){
					pq.enqueue(Point{x:(pos.x - 1), y:pos.y})
					fmt.Println("added something :) left")
				}
			}
		}
	}

	fmt.Println(sum)
}