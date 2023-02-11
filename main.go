package main

import "fmt"

func main() {
	fmt.Println("Starting...")

	grid := NewGrid()
	grid2 := NewGrid()

	grid.Print()
	grid2.Print()

	fmt.Println("Terminating...")
}

type Digit struct {
	Value int
}

type Box struct {
	Value [][]Digit
}

type Grid struct {
	Value [][]Box
}

func NewGrid() *Grid {
	return &Grid{
		Value: [][]Box{},
	}
}

func (g *Grid) Print() {
	fmt.Println("Printing")
}

func (g *Grid) GetRow(rowIndex int) []Digit {
	gridIndex := rowIndex / 3
	boxIndex := rowIndex % 3
	var row []Digit
	for _, box := range g.Value[gridIndex] {
		row = append(row, box.Value[boxIndex]...)
	}
	return row
}
