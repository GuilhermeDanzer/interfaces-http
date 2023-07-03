package main

import "fmt"

type shapes interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	side float64
}

func main() {

	sq := square{side: 5}
	tr := triangle{height: 15, base: 4}
	printArea(sq)
	printArea(tr)
}

func printArea(s shapes) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {

	return t.base * t.height * 0.5
}

func (sq square) getArea() float64 {

	return sq.side * 2
}
