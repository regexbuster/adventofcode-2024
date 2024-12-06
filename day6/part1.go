package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"slices"
)

type Direction int

const (
	Left	Direction = iota
	Up
	Right
	Down
)

var directionName = map[Direction]string{
    Left:	"left",
    Up: 	"up",
    Right:	"right",
    Down:	"down",
}

type TileType int

const (
	TileEmpty	TileType = iota
	TileWall
	TileEnd
	TileError
)

type Guard struct {
	X 		int
	Y 		int
	Facing	Direction
	Visited	map[string]bool
}

func (g Guard) Print() {
	fmt.Printf("At (%d, %d) and facing %s\n", g.X, g.Y, directionName[g.Facing])
}

func (g *Guard) TurnRight() {
	switch g.Facing {
	case Left:
		g.Facing = Up
	case Up:
		g.Facing = Right
	case Right:
		g.Facing = Down
	case Down:
		g.Facing = Left
	default:
		panic("CANNOT TURN RIGHT ABORT!!!")
	}
}

func (g *Guard) StepForward() {
	switch g.Facing {
	case Left:
		g.X = g.X - 1
	case Up:
		g.Y = g.Y - 1
	case Right:
		g.X = g.X + 1
	case Down:
		g.Y = g.Y + 1
	default:
		panic("CANNOT MOVE FORWARD ABORT!!!")
	}

	g.Visited[fmt.Sprintf("%d,%d", g.X, g.Y)] = true
}

type Board [][]string

func (b Board) Print() {
	for _, row := range b {
		fmt.Println(row)
	}
}

func (b Board) whichTile(where Direction, x int, y int) TileType {
	switch where {
	case Left:
		if x - 1 >= 0{
			if b[y][x-1] == "#" {
				return TileWall
			}

			return TileEmpty
		}
		return TileEnd
	case Up:
		if y - 1 >= 0{
			if b[y-1][x] == "#" {
				return TileWall
			}

			return TileEmpty
		}
		return TileEnd
	case Right:
		if x + 1 < len(b[y]){
			if b[y][x+1] == "#" {
				return TileWall
			}

			return TileEmpty
		}
		return TileEnd
	case Down:
		if y + 1 < len(b){
			if b[y+1][x] == "#" {
				return TileWall
			}

			return TileEmpty
		}
		return TileEnd
	default:
		return TileError
	}
} 

func GuardView(b Board, g Guard) {
	for rowNum, row := range b {
		if rowNum == g.Y {
			tmpRow := make([]string, len(row))
			copy(tmpRow, row)
			

			switch g.Facing {
			case Left:
				tmpRow[g.X] = "<"
			case Up:
				tmpRow[g.X] = "^"
			case Right:
				tmpRow[g.X] = ">"
			case Down:
				tmpRow[g.X] = "v"
			}

			fmt.Println(strings.Join(tmpRow, ""))
		} else {
			fmt.Println(strings.Join(row, ""))
		}
	}
}

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

	var board Board
	var guard Guard

	var rowCount int

	for s.Scan(){
		text := s.Text()

		row := strings.Split(text, "")

		index := slices.Index(row, "^")

		if index != -1 {
			guard = Guard{X:index, Y:rowCount, Facing:Up, Visited: make(map[string]bool)}
			row[index] = "."
		}

		board = append(board, row)

		rowCount += 1
	}

	GuardView(board, guard)
	guard.Print()

	fmt.Println("\n\n\n\n")

	for i := 0; i < 100000; i++ {
		facingTile := board.whichTile(guard.Facing, guard.X, guard.Y)

		switch facingTile {
		case TileEmpty:
			guard.StepForward()
		case TileWall:
			guard.TurnRight()
		case TileEnd:
			fmt.Println("Ended b/c hit edge")
			i = 1000000
		case TileError:
			panic("Tile error")
		default:
			panic("Something wrong with handling next tile")
		}
	}

	GuardView(board, guard)
	guard.Print()

	var sum int

	for _, v := range guard.Visited {
		if v {
			sum += 1
		}
	}

	fmt.Printf("Total Visited: %d\n", sum)
}