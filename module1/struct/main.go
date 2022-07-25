package main

import "fmt"

type human struct {
	Weight float64
	Height float64
}

func main() {
	h := human{Weight: 68.9, Height: 175.0}
	s := make([]float64, 0)

	fmt.Println(getWeight(&h, s))
}

func getWeight(h *human, s []float64) []float64 {
	return append(s, h.Weight, h.Height)
}
