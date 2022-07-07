package main

import (
	"fmt"
)

type TwoNumbers struct {
	X, Y int
}

func (v TwoNumbers) AddNumbers() int {
	return v.X + v.Y
}

func (v TwoNumbers) MulNumbers() int {
	return v.X * v.Y
}

func main() {
	v := TwoNumbers{3, 4}

	fmt.Println("Sum =", v.AddNumbers())
	fmt.Println("Mul =", v.MulNumbers())
}
