package main

import (
	"fmt"
	"math/rand"
)

func runTrial(switchBoxes bool, nBoxes int) bool {

	chosenDoor := rand.Intn(nBoxes) + 1

	if switchBoxes {
		revealedDoor := 3
		if chosenDoor == 2 {
			revealedDoor = 3
		} else {
			revealedDoor = 2
		}

		var availableDoors []int
		for dnum := 1; dnum <= nBoxes; dnum++ {
			if dnum != chosenDoor && dnum != revealedDoor {
				availableDoors = append(availableDoors, dnum)
			}
		}

		chosenDoor = availableDoors[rand.Intn(len(availableDoors))]
	}

	return chosenDoor == 1
}

func runTrials(nTrials int, switchDoors bool, nDoors int) int {
	nWins := 0

	for i := 0; i < nTrials; i++ {
		if runTrial(switchDoors, nDoors) {
			nWins++
		}
	}

	return nWins
}

func main() {
	nDoors, nTrials := 3, 10000
	nWinsWithoutSwitch := runTrials(nTrials, false, nDoors)
	nWinsWithSwitch := runTrials(nTrials, true, nDoors)

	fmt.Printf("Monty Hall Problem with %d doors\n", nDoors)
	fmt.Printf("Proportion of wins without switching: %.4f\n", float64(nWinsWithoutSwitch)/float64(nTrials))
	fmt.Printf("Proportion of wins with switching: %.4f\n", float64(nWinsWithSwitch)/float64(nTrials))
}

// package main

// import "fmt"
// import "math/rand"


// func main() {

// 	no_rounds := 10
// 	var player_strat bool
// 	var winning_rounds int32

// 	for i:=0; i < no_rounds; i++ {

// 		winning_number:=rand.Intn(3)
// 		player_choice:=rand.Intn(3)
// 		choices := make([]bool, 3)
// 		choices[winning_number] = true
// 		if player_strat == false && choices[player_choice] == true {
// 			winning_rounds++
// 		}
// 	}

//     fmt.Println("Hello, World!")
// 	fmt.Println(winning_rounds)
// }