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
	value [][]Digit
}

/*
*
TODO: Write unite tests for all GetRow and Col methods
*/
func (b *Box) GetRow(rowIndex int) []Digit {
	return b.value[rowIndex]
}

func (b *Box) AddRow(rowIndex int, rowData []Digit) {
	b.value[rowIndex] = rowData
}

func (b *Box) GetCol(colIndex int) []Digit {
	var col []Digit
	for _, row := range b.value {
		col = append(col, row[colIndex])
	}
	return col
}

type Grid struct {
	value [][]Box
}

/*
*
TODO: Accept data for constructor
*/
func NewGrid() *Grid {
	return &Grid{
		value: [][]Box{},
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
	for _, box := range g.value[gridIndex] {
		row = append(row, box.GetRow(boxIndex)...)
	}
	return row
}

func (g *Grid) AddRow(rowIndex int, rowData []Digit) {
	gridIndex := rowIndex / 3
	boxIndex := rowIndex % 3

	gridRow := g.value[gridIndex]

	for i, box := range gridRow {
		low := i * 3
		high := low + 3

		box.AddRow(boxIndex, rowData[low:high])
	}

}

func (g *Grid) GetCol(colIndex int) []Digit {
	boxIndex := colIndex % 3
	var col []Digit
	for _, rowBoxes := range g.value {
		col = append(col, rowBoxes[boxIndex].GetCol(boxIndex)...)
	}
	return col
}
