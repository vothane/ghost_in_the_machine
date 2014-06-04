package main

import (
	"fmt"
	"math"
	"sort"
)

func sumCalc(coll []float64, summer func(x float64) float64) float64 {
	var sum float64 = 0.0

	for _, x := range coll {
		sum += summer(x)
	}
	return sum
}

func sumOfProds(x, y []float64) float64 {
	var sum float64 = 0.0

	for i := range x {
		sum += x[i] * y[i]
	}
	return sum
}

func pearson(v1, v2 []float64) float64 {
	sum1 := sumCalc(v1, func(x float64) float64 { return x })
	sum2 := sumCalc(v2, func(x float64) float64 { return x })

	sum1sq := sumCalc(v1, func(x float64) float64 { return x * x })
	sum2sq := sumCalc(v2, func(x float64) float64 { return x * x })

	psum := sumOfProds(v1, v2)

	num := psum - (sum1 * sum2 / float64(len(v1)))
	den := math.Sqrt((sum1sq - math.Pow(sum1, 2)/float64(len(v1))) * (sum2sq - math.Pow(sum2, 2)/float64(len(v1))))

	if den == 0 {
		return 0.0
	}
	return 1.0 - num/den
}

func colMinMax(data [][]float64) [2][2]float64 {
	var colData [2][7]float64
	var minmax [2][2]float64

	for i, point := range data {
		for j, item := range point {
			colData[j][i] = item
		}
	}

	for i := range colData {
		sort.Float64s((colData[i])[:])
	}

	for i, datum := range colData {
		minmax[i] = [...]float64{datum[0], datum[6]}
	}
	return minmax
}

func main() {
	v1 := []float64{0.0, 1.0, 0.0, 0.0, 3.0, 3.0, 0.0, 0.0, 3.0, 0.0, 6.0, 0.0, 1.0, 0.0, 4.0, 3.0, 0.0, 0.0, 0.0, 0.0, 0.0, 4.0, 0.0}
	v2 := []float64{0.0, 2.0, 1.0, 0.0, 6.0, 2.0, 1.0, 0.0, 4.0, 5.0, 25.0, 0.0, 0.0, 0.0, 6.0, 12.0, 4.0, 2.0, 1.0, 4.0, 0.0, 3.0, 0.0}
	rows := [][]float64{{1.0, 1.0}, {1.5, 2.0}, {3.0, 4.0}, {5.0, 7.0}, {3.5, 5.0}, {4.5, 5.0}, {3.5, 4.5}}
	fmt.Println(pearson(v1, v2))
	fmt.Println(rows)
	fmt.Println(colMinMax(rows))
}
