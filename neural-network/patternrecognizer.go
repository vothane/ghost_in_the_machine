package main

import (
	"fmt"
	"math/rand"
	"time"
	//matrix "github.com/skelterjohn/go.matrix"
)

func activate(sum int) int {
	if sum > 0 {
		return 1
	} else {
		return -1
	}
  return 1
}

type PatternRecognizer struct{ weights []int }

func New(n int) *PatternRecognizer {
	ws := make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := range ws {
		if r := rand.Intn(2); r == 0 {
			ws[i] = -1
		} else {
			ws[i] = 1
		}
	}
	patternrecog := PatternRecognizer{weights: ws}
	return &patternrecog
}

func (patternrecog *PatternRecognizer) Feedforward(inputs []int) int {
	sum := 0
	for i := range patternrecog.weights {
		sum += inputs[i] * patternrecog.weights[i]
	}
	return activate(sum)
}

func main() {
	patternrecog := New(3)
  point := []int{50,-12,1};
  result := patternrecog.Feedforward(point);
	fmt.Println(result)
}
