package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"slices"
)

type Point struct {
	x	int
	y	int
}

func addPoints(a, b Point) Point {
	return Point{x:(a.x + b.x), y:(a.y + b.y)}
}

func subPoints(a, b Point) Point {
	return Point{x:(a.x - b.x), y:(a.y - b.y)}
}

func samePoint(a, b Point) bool {
	return a.x == b.x && a.y == b.y
}

type AntennaNetwork struct {
	name				string
	antennas			[]Point
	possibleAntinodes	[]Point
}

func createNetwork(name string, pos Point) AntennaNetwork{
	return AntennaNetwork{name: name, antennas: []Point{pos}, possibleAntinodes: []Point{}}
}

func (an *AntennaNetwork) updateAntinodes(xBound, yBound int) {
	for i := 0; i < len(an.antennas); i++ {
		for j := i + 1; j < len(an.antennas); j++ {
			deltaPoint := subPoints(an.antennas[i], an.antennas[j])

			an.possibleAntinodes = append(an.possibleAntinodes, an.antennas[i])

			for i := addPoints(an.antennas[i], deltaPoint); i.x >= 0 && i.x < xBound && i.y >= 0 && i.y < yBound; i = addPoints(i, deltaPoint) {
				if !slices.Contains(an.possibleAntinodes, i) {
					an.possibleAntinodes = append(an.possibleAntinodes, i)
				}
			}

			for i := subPoints(an.antennas[i], deltaPoint); i.x >= 0 && i.x < xBound && i.y >= 0 && i.y < yBound; i = subPoints(i, deltaPoint) {
				if !slices.Contains(an.possibleAntinodes, i) {
					an.possibleAntinodes = append(an.possibleAntinodes, i)
				}
			}
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./puzzleInput.txt")
	check(err)

	s := bufio.NewScanner(file)

	antennaMap := make(map[string]AntennaNetwork)

	var xMax int
	var yIndex int

	for s.Scan() {
		text := s.Text()
		stext := strings.Split(text, "")
		xMax = len(stext)

		for xIndex, value := range stext {
			if value == "."{
				continue
			}

			data, ok := antennaMap[value]

			if ok {
				data.antennas = append(data.antennas, Point{xIndex, yIndex})
				antennaMap[value] = data
			} else {
				antennaMap[value] = createNetwork(value, Point{xIndex, yIndex})
			}
		}

		yIndex = yIndex + 1
	}

	var finalAntinodes []Point

	for i, v := range antennaMap {
		v.updateAntinodes(xMax, yIndex)
		antennaMap[i] = v

		// finalAntinodes = slices.Concat(finalAntinodes, v.possibleAntinodes)

		for _, val := range v.possibleAntinodes {
			if !slices.Contains(finalAntinodes, val){
				finalAntinodes = append(finalAntinodes, val)
			}
		}

		fmt.Println(v)
	}

	fmt.Println(xMax, yIndex)
	fmt.Println(finalAntinodes, len(finalAntinodes))
}