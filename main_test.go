package main

import (
	"fmt"
	"testing"
)

func TestBoxGetRow(t *testing.T) {

	expected := []Digit{{Value: 1}, {Value: 2}, {Value: 3}}

	box := NewBox()
	fmt.Println(box)
	box.AddRow(0, expected)

	print(box.GetRow(0))

	output := box.GetRow(0)

	println("hello")

	if (*[3]Digit)(output) != (*[3]Digit)(expected) {
		t.Errorf("got %v, wanted %v", output, expected)
	}
}
