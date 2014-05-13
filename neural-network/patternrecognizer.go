package main

import (
	"fmt"
  "math/rand"
  "time"
  //"testing"
	//matrix "github.com/skelterjohn/go.matrix"
)

type PatternRecognizer struct{ weights []int }

func New(n int) *PatternRecognizer{
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

func main() {
  patternrecog := New(5)
  fmt.Println(patternrecog)
}

