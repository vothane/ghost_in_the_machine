package main

import (
	"fmt"
)

func simdist(prefs map[string]map[string]float64, person1 string, person2 string) map[string]int {
	si := make(map[string]int)
	
	for key := range prefs[person1] {
		if _, ok := prefs[person2][key]; ok {
			si[key] = 1
		}
	}
	return si
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
	result := simdist(critics, "Lisa Rose", "Gene Seymour")
	fmt.Println(result)
}
