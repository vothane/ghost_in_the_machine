package main

import (
	"fmt"
	"math"
	"sort"
)

func simdist(prefs map[string]map[string]float64, person1 string, person2 string) float64 {
	si := make(map[string]int)
	for key := range prefs[person1] {
		if _, ok := prefs[person2][key]; ok {
			si[key] = 1
		}
	}

	sumofsquares := 0.0
	for key := range prefs[person1] {
		if _, ok := prefs[person2][key]; ok {
			sumofsquares += math.Pow((prefs[person1][key] - prefs[person2][key]), 2)
		}
	}
	return 1 / (1 + sumofsquares)
}

func simpearson(prefs map[string]map[string]float64, person1 string, person2 string) float64 {
	si := make(map[string]int)
	for key := range prefs[person1] {
		if _, ok := prefs[person2][key]; ok {
			si[key] = 1
		}
	}

	sum1 := 0.0
	sum2 := 0.0
	sum1Sq := 0.0
	sum2Sq := 0.0
	pSum := 0.0
	for key := range prefs[person1] {
		if _, ok := prefs[person2][key]; ok {
			sum1 += prefs[person1][key]
			sum1Sq += math.Pow(prefs[person1][key], 2)
		}
	}
	for key := range prefs[person2] {
		if _, ok := prefs[person1][key]; ok {
			sum2 += prefs[person2][key]
			sum2Sq += math.Pow(prefs[person2][key], 2)
			pSum += prefs[person1][key] * prefs[person2][key]
		}
	}
	n := float64(len(si))
	num := pSum - (sum1 * sum2 / n)
	den := math.Sqrt((sum1Sq - math.Pow(sum1, 2)/n) * (sum2Sq - math.Pow(sum2, 2)/n))
	r := num / den
	return r
}

type Rank struct {
	Score float64
	Name  string
}

type Ranks []Rank

func (r Ranks) Len() int           { return len(r) }
func (r Ranks) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Ranks) Less(i, j int) bool { return r[i].Score > r[j].Score }

func lookup(m map[string]float64, k string) float64 {
	if val, ok := m[k]; ok {
		return val
	}
	return 0.0
}

func getrecommend(prefs map[string]map[string]float64, person string, similarity func(a map[string]map[string]float64, b string, c string) float64) []Rank {
	totals := make(map[string]float64)
	simsums := make(map[string]float64)
	rankings := make([]Rank, 0, 6)

	for other := range prefs {
		if other == person {
			continue
		}
		sim := similarity(prefs, person, other)
		if sim == 0 {
			continue
		}
		for movie := range prefs[other] {
			_, ok := prefs[other][movie]
			score := prefs[person][movie]
			if ok != true || score == 0 {
				totals[movie] += prefs[other][movie] * sim
				simsums[movie] += sim
			}
		}
	}
	for movie := range totals {
		rankings = append(rankings, Rank{totals[movie] / simsums[movie], movie})
	}
	sort.Sort(Ranks(rankings))
	return rankings
}

func main() {
	critics := map[string]map[string]float64{
		"Lisa Rose": {"Lady in the Water": 2.5, "Snakes on a Plane": 3.5,
			"Just My Luck": 3.0, "Superman Returns": 3.5, "You, Me and Dupree": 2.5,
			"The Night Listener": 3.0},
		"Gene Seymour": {"Lady in the Water": 3.0, "Snakes on a Plane": 3.5,
			"Just My Luck": 1.5, "Superman Returns": 5.0, "The Night Listener": 3.0,
			"You, Me and Dupree": 3.5},
		"Michael Phillips": {"Lady in the Water": 2.5, "Snakes on a Plane": 3.0,
			"Superman Returns": 3.5, "The Night Listener": 4.0},
		"Claudia Puig": {"Snakes on a Plane": 3.5, "Just My Luck": 3.0,
			"The Night Listener": 4.5, "Superman Returns": 4.0,
			"You, Me and Dupree": 2.5},
		"Mick LaSalle": {"Lady in the Water": 3.0, "Snakes on a Plane": 4.0,
			"Just My Luck": 2.0, "Superman Returns": 3.0, "The Night Listener": 3.0,
			"You, Me and Dupree": 2.0},
		"Jack Matthews": {"Lady in the Water": 3.0, "Snakes on a Plane": 4.0,
			"The Night Listener": 3.0, "Superman Returns": 5.0, "You, Me and Dupree": 3.5},
		"Toby": {"Snakes on a Plane": 4.5, "You, Me and Dupree": 1.0, "Superman Returns": 4.0},
	}
	resultsimdist := simdist(critics, "Lisa Rose", "Gene Seymour")
	resultsimpearson := simpearson(critics, "Lisa Rose", "Gene Seymour")
	recommendation := getrecommend(critics, "Toby", simpearson)
	fmt.Println(resultsimdist)
	fmt.Println(resultsimpearson)
	fmt.Println(recommendation)
}
