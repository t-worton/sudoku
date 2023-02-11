package main

import "fmt"

func main() {
	fmt.Println("Starting...")
	grid := NewGrid()
	grid.Print()
	fmt.Println("Terminating...")
}

type Digit struct {
	Value int
}

type Box struct {
	Value [][]Digit
}

/*
*
TODO: Write unite tests for all GetRow and Col methods
*/
func (b *Box) GetRow(rowIndex int) []Digit {
	return b.Value[rowIndex]
}

func (b *Box) GetCol(colIndex int) []Digit {
	var col []Digit
	for _, row := range b.Value {
		col = append(col, row[colIndex])
	}
	return col
}

type Grid struct {
	Value [][]Box
}

/*
*
TODO: Accept data for constructor
*/
func NewGrid() *Grid {
	return &Grid{
		Value: [][]Box{},
	}
}

func (g *Grid) Print() {
	fmt.Println("Printing")
}

/*
*
TODO: Write unite tests for all GetRow and Col methods
*/
func (g *Grid) GetRow(rowIndex int) []Digit {
	gridIndex := rowIndex / 3
	boxIndex := rowIndex % 3
	var row []Digit
	for _, box := range g.Value[gridIndex] {
		row = append(row, box.GetRow(boxIndex)...)
	}
	return row
}

func (g *Grid) GetCol(colIndex int) []Digit {
	boxIndex := colIndex % 3
	var col []Digit
	for _, rowBoxes := range g.Value {
		col = append(col, rowBoxes[boxIndex].GetCol(boxIndex)...)
	}
	return col
}
