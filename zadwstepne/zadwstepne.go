package main

import (
	"fmt"
	"math/rand"
)

func pojedynczaProba(zamianaPudelek bool, nPudelek int) bool {

	wybranePudelko := rand.Intn(nPudelek) + 1

	if zamianaPudelek {
		odkrytePudelko := 3
		if wybranePudelko == 2 {
			odkrytePudelko = 3
		} else {
			odkrytePudelko = 2
		}

		var dostepnePudelka []int
		for i := 1; i <= nPudelek; i++ {
			if i != wybranePudelko && i != odkrytePudelko {
				dostepnePudelka = append(dostepnePudelka, i)
			}
		}

		wybranePudelko = dostepnePudelka[rand.Intn(len(dostepnePudelka))]
	}

	return wybranePudelko == 1
}

func generatorProb(iloscProb int, zamienPudelka bool, nPudelek int) int {
	nWygrane := 0

	for i := 0; i < iloscProb; i++ {
		if pojedynczaProba(zamienPudelka, nPudelek) {
			nWygrane++
		}
	}

	return nWygrane
}

func main() {
	nPudelek, nProb := 3, 10000
	nWygraneBezZamiany := generatorProb(nProb, false, nPudelek)
	nWygraneZZamiana := generatorProb(nProb, true, nPudelek)

	fmt.Printf("ilosc pudelek: %d ilosc prob: %d\n", nPudelek, nProb)
	fmt.Printf("procent wygranych bez zamiany pudelek: %.4f \n", float64(nWygraneBezZamiany)/float64(nProb))
	fmt.Printf("procent wygranych z zamiana pudelek: %.4f \n", float64(nWygraneZZamiana)/float64(nProb))
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