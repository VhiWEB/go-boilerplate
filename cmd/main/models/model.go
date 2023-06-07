package model

import (
	"fmt"
	"math"
)

type Model interface {
	get() float64
	keliling() float64
}

func init() {
	fmt.Print(math.Sqrt(12))
}
