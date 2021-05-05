package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64  {
	return math.Pow(p.sisi, 2)
}

func main() {
	var h hitung
	h = persegi{sisi:10.0}
	fmt.Println("Luas Persegi:", h.luas())
}
