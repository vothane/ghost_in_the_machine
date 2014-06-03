package main

import (
	"fmt"
	"math"
)

func sumcalc(coll []float64, summer func(x float64) float64) float64 {
	var sum float64 = 0.0

	for _, x := range coll {
		sum += summer(x)
	}
	return sum
}

func sumofprods(x, y []float64) float64 {
	var sum float64 = 0.0

	for i := range x {
		sum += x[i] * y[i]
	}
	return sum
}

func pearson(v1, v2 []float64) float64 {
	sum1 := sumcalc(v1, func(x float64) float64 { return x })
	sum2 := sumcalc(v2, func(x float64) float64 { return x })

	sum1sq := sumcalc(v1, func(x float64) float64 { return x * x })
	sum2sq := sumcalc(v2, func(x float64) float64 { return x * x })

	psum := sumofprods(v1, v2)

	num := psum - (sum1 * sum2 / float64(len(v1)))
	den := math.Sqrt((sum1sq - math.Pow(sum1, 2) / float64(len(v1))) * (sum2sq-math.Pow(sum2, 2) / float64(len(v1))))

	if den == 0 {
		return 0.0
	}

	return 1.0 - num / den
}

func main() {
	v1 := []float64{0.0, 1.0, 0.0, 0.0, 3.0, 3.0, 0.0, 0.0, 3.0, 0.0, 6.0, 0.0, 1.0, 0.0, 4.0, 3.0, 0.0, 0.0, 0.0, 0.0, 0.0, 4.0, 0.0}
	v2 := []float64{0.0, 2.0, 1.0, 0.0, 6.0, 2.0, 1.0, 0.0, 4.0, 5.0, 25.0, 0.0, 0.0, 0.0, 6.0, 12.0, 4.0, 2.0, 1.0, 4.0, 0.0, 3.0, 0.0}
	fmt.Println(pearson(v1, v2))
}
