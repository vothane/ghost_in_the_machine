package main

import (
	"fmt"
	"math/rand"
	"time"
	//matrix "github.com/skelterjohn/go.matrix"
)

type PatternRecognizer struct{ weights []float64 }

func New(n int) *PatternRecognizer {
	ws := make([]float64, n)
	rand.Seed(time.Now().Unix())
	for i := range ws {
		if r := rand.Intn(1); r == 0 {
			ws[i] = float64(-1)
		} else {
			ws[i] = float64(1)
		}
	}
	patternrecog := PatternRecognizer{weights: ws}
	return &patternrecog
}

func (patternrecog PatternRecognizer) Feedforward(inputs []float64) float64 {
	sum := 0.0
	for i := range patternrecog.weights {
		sum += inputs[i] * patternrecog.weights[i]
	}
	return patternrecog.activate(sum)
}

func (patternrecog*PatternRecognizer) activate(sum float64) float64 {
	var val float64
	if sum > 0 {
		val = 1.0
	} else {
		val = -1.0
	}
	return val
}

func (patternrecog PatternRecognizer) Train(inputs []float64, desired float64) {
	guess := patternrecog.Feedforward(inputs)
	error := desired - guess
	for i := range patternrecog.weights {
		patternrecog.weights[i] += 0.01 * error * inputs[i]
	}
}

func main() {
	patternrecog := New(3)
	point := []float64{50, -12, 1}
	result := patternrecog.Feedforward(point)
	fmt.Println(result)
}
