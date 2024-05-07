package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
    fireResistance int
    isTree         bool
}

//utworzenie lasu bez dodatkowego parametru resistance
func initForestNoFireResistance(size int, treeProbability float64) [][]Tree {
    forest := make([][]Tree, size)
    for i := range forest {
        forest[i] = make([]Tree, size)
        for j := range forest[i] {
            if rand.Float64() < treeProbability {
                forest[i][j] = Tree{0, true}  // wszystkie drzewa maja resistance ustawiony na 0(odrazu rozprzestrzeniaja ogien)
            } else {
                forest[i][j] = Tree{0, false} // brak drzewa
            }
        }
    }
    return forest
}

//prosty model spalania, bez dodatkowego parametru
func simpleSpreadFire(forest [][]Tree, x int, y int) {
    if x < 0 || y < 0 || x >= len(forest) || y >= len(forest) || !forest[x][y].isTree {
        return
    }
    //  Oznaczenie biezacego drzewa jako spalone i proste rozprzestrzenianie ognia na wszystkie sssiednie drzewa
    forest[x][y].isTree = false
    forest[x][y].fireResistance = -1 // oznaczenie drzewa jako spalonego

	//dx, dy, nx, ny kierunki rozsprzestrzeniania
    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx != 0 || dy != 0 {
                nx, ny := x + dx, y + dy
                if nx >= 0 && ny >= 0 && nx < len(forest) && ny < len(forest) && forest[nx][ny].isTree {
                    simpleSpreadFire(forest, nx, ny)
                }
            }
        }
    }
}

//inicjacja lasu z dodatkowym losowym parametrem resistance
func initForestWithResistance(size int, treeProbability float64) [][]Tree {
    forest := make([][]Tree, size)
    for i := range forest {
        forest[i] = make([]Tree, size)
        for j := range forest[i] {
            if rand.Float64() < treeProbability {
                fireResistance := rand.Intn(2) // Randomly 0 or 1
                forest[i][j] = Tree{fireResistance, true}
            } else {
                forest[i][j] = Tree{0, false}
            }
        }
    }
    return forest
}

//model spalanai uwzgledniajacy resistance na ogien drzew
//w przypadku gdy resistance = 1 to drzewo sie nie spali ani nie rozprzestrzeni ognia
//ale jego resistance zostanie zmniejszony do 0
func spreadFireWithTreeResistance(forest [][]Tree, x int, y int) {
    if x < 0 || y < 0 || x >= len(forest) || y >= len(forest) || !forest[x][y].isTree {
        return
    }
    if forest[x][y].fireResistance == 0 {
        forest[x][y].isTree = false
        forest[x][y].fireResistance = -1 // oznaczenie drzewa jako spalone
		//kierunki rozprzestrzeniania, tym razem w innej formie danych
        directions := [][2]int{
            {-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
        }
        for _, d := range directions {
            nx, ny := x+d[0], y+d[1]
            if nx >= 0 && ny >= 0 && nx < len(forest) && ny < len(forest) {
                if forest[nx][ny].isTree {
                    if forest[nx][ny].fireResistance == 1 {
                        //zmniejszenie odpornosci na ogien w przypadku napotkania ognia
                        forest[nx][ny].fireResistance = 0
                    } else {
                        spreadFireWithTreeResistance(forest, nx, ny)
                    }
                }
            }
        }
    }
}

//funkcja do zliczania ilosci spalonych drzew
func countBurntTrees(forest [][]Tree) (burnt int) {
	for i := range forest {
		for j := range forest[i] {
			if !forest[i][j].isTree && forest[i][j].fireResistance == -1 {
				burnt++
			}
		}
	}
	return
}

//funkcja do zliczania pustych miejsc
func countEmptyLots(forest [][]Tree) (empty int) {
	for i := range forest {
		for j := range forest[i] {
			if !forest[i][j].isTree && forest[i][j].fireResistance != -1 {
				empty++
			}
		}
	}
	return
}

//funkcja do przedstawienia graficznego lasu na konsoli
func printForest(forest [][]Tree) {
	for _, row := range forest {
		for _, cell := range row {
			if cell.isTree {
				fmt.Print("T ")
			} else if !cell.isTree && cell.fireResistance == -1 {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {

	size := 10
	totalTrials := 1000
	trials := 100
	minDensity := 0.45
	maxDensity := 0.96
	step := 0.025

	densityWinCount := make(map[float64]int)
	densityWinCountNoWindDirection := make(map[float64]int)

	//proby dla modelu spalania lasu z odpornoscia
	for t := 0; t < totalTrials; t++ {
		bestDensity := 0.0
		lowestBurnPercentage := 100.0

		for density := minDensity; density <= maxDensity; density += step {
			var totalBurntPercentage float64

			for i := 0; i < trials; i++ {
				forest := initForestWithResistance(size, density)
				startX, startY := rand.Intn(size), rand.Intn(size)
				spreadFireWithTreeResistance(forest, startX, startY)
				burntTrees := countBurntTrees(forest)
				emptySpaces := countEmptyLots(forest)
				burntPercentage := 100 * float64(burntTrees) / (float64(size*size) - float64(emptySpaces))
				totalBurntPercentage += burntPercentage
			}

			averageBurntPercentage := totalBurntPercentage / float64(trials)

			if averageBurntPercentage < lowestBurnPercentage {
				lowestBurnPercentage = averageBurntPercentage
				bestDensity = density
			}
		}

		densityWinCount[bestDensity]++
	}

	//proby dla prostego modelu spalania lasu
	for t := 0; t < totalTrials; t++ {
		bestDensity := 0.0
		lowestBurnPercentage := 100.0

		for density := minDensity; density <= maxDensity; density += step {
			var totalBurntPercentage float64

			for i := 0; i < trials; i++ {
				forest := initForestNoFireResistance(size, density)
				startX, startY := rand.Intn(size), rand.Intn(size)
				simpleSpreadFire(forest, startX, startY)
				burntTrees := countBurntTrees(forest)
				emptySpaces := countEmptyLots(forest)
				burntPercentage := 100 * float64(burntTrees) / (float64(size*size) - float64(emptySpaces))
				totalBurntPercentage += burntPercentage
			}

			averageBurntPercentage := totalBurntPercentage / float64(trials)

			if averageBurntPercentage < lowestBurnPercentage {
				lowestBurnPercentage = averageBurntPercentage
				bestDensity = density
			}
		}

		densityWinCountNoWindDirection[bestDensity]++
	}

	fmt.Println("Wyniki dla modelu spalania z odpornoscia na ogien, ilosc prob:", totalTrials)
	for density, count := range densityWinCount {
		fmt.Printf("Gestosc zalesienia %.2f pojawila sie %d razy jako ta z najnizszym stopniem wypalenia.\n", density, count)
	}
	fmt.Println("Wyniki dla prostego modelu spalania, ilosc prob:", totalTrials)
	for density, count := range densityWinCountNoWindDirection {
		fmt.Printf("Gestosc zalesienia %.2f pojawila sie %d razy jako ta z najnizszym stopniem wypalenia.\n", density, count)
	}
	fmt.Println("_______________________")

	density := 0.3 + rand.Float64() * (0.8 - 0.3)
	forest := initForestNoFireResistance(size, density)
	printForest(forest)
	startX, startY := rand.Intn(size), rand.Intn(size)
	simpleSpreadFire(forest, startX, startY)
	burntTrees := countBurntTrees(forest)
	emptySpaces := countEmptyLots(forest)
	println("______________________________")
	printForest(forest)
	fmt.Printf("Procent spalonego lasu z gestoscia zalesienia %.2f: %.2f%%\n", density, 100*float64(burntTrees)/(float64(size*size)-float64(emptySpaces)))
}
