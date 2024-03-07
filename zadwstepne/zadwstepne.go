package main

import "fmt"
import "math/rand"


func main() {

	no_rounds := 10
	var player_strat bool
	var winning_rounds int32

	for i:=0; i < no_rounds; i++ {

		winning_number:=rand.Intn(3)
		player_choice:=rand.Intn(3)
		choices := make([]bool, 3)
		choices[winning_number] = true
		if player_strat == false && choices[player_choice] == true {
			winning_rounds++
		}
	}

    fmt.Println("Hello, World!")
	fmt.Println(winning_rounds)
}