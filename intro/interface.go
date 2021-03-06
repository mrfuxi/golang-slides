package main

import (
	"fmt"
	"math"
)

// START OMIT
type ShapeArea interface { // HL
	Area() float64 // HL
} // HL

type Square struct{ Side float64 }

func (s Square) Area() float64  { return s.Side * s.Side } // HL
func (s Square) String() string { return fmt.Sprintf("<Square %v>", s.Side) }

type Circle struct{ Radius float64 }

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius } // HL

// MID OMIT
func SumAreas(shape ...ShapeArea) (sum float64) { // HL
	for _, f := range shape {
		sum += f.Area() // HL
	}
	return
}

func main() {
	sqare1 := Square{Side: 1}
	sqare2 := Square{Side: 2}
	circle := Circle{Radius: 3}
	sum := SumAreas(sqare1, sqare2, circle)
	fmt.Printf("%v + %v + %v = %v\n", sqare1, sqare2, circle, sum)
}

// END OMIT
