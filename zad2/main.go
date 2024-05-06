package main

import (
	"fmt"
	"math/rand"
	//"time"
)

const (
	Empty = 0
	Tree  = 1
	Burnt = 2
)

var windDirection = [2]int{1, 1} // Southeast wind, modify as needed

func initForest(size int, treeProbability float64) [][]int {
	forest := make([][]int, size)
	for i := range forest {
		forest[i] = make([]int, size)
		for j := range forest[i] {
			if rand.Float64() < treeProbability {
				forest[i][j] = Tree
			}
		}
	}
	return forest
}

func spreadFire(forest [][]int, x int, y int) {
	if x < 0 || y < 0 || x >= len(forest) || y >= len(forest) || forest[x][y] != Tree {
		return
	}
	forest[x][y] = Burnt
	// Apply wind influence by reordering spread directions based on wind
	directions := [][2]int{
		{windDirection[0], windDirection[1]}, // wind primary direction
		{windDirection[0], 0}, {0, windDirection[1]}, // secondary influenced directions
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}, // remaining directions
	}

	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < len(forest) && ny < len(forest) && forest[nx][ny] == Tree {
			spreadFire(forest, nx, ny)
		}
	}
}

func countBurntTrees(forest [][]int) (burnt int) {
	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == Burnt {
				burnt++
			}
		}
	}
	return
}

func countEmptyLots(forest [][]int) (empty int) {
	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == Empty {
				empty++
			}
		}
	}
	return
}

func printForest(forest [][]int) {
	for _, row := range forest {
		for _, cell := range row {
			if cell == Tree {
				fmt.Print("T ")
			} else if cell == Burnt {
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
	trials := 100
	minDensity := 0.1
	maxDensity := 0.91
	step := 0.05
	bestBurntPercentage := 100.0
	optimalDensity := 0.0

	for density := minDensity; density <= maxDensity; density += step {
		var totalBurntPercentage float64
		for i := 0; i < trials; i++ {
			forest := initForest(size, density)
			if (i == 0) {
				printForest(forest)
				println("__________________")
			}
			startX, startY := rand.Intn(size), rand.Intn(size)
			spreadFire(forest, startX, startY)
			burntTrees := countBurntTrees(forest)
			emptySpaces := countEmptyLots(forest)
			if (i == 0) {
				printForest(forest)
			}
			totalBurntPercentage += 100 * float64(burntTrees) / (float64(size*size) - float64(emptySpaces))
			if (totalBurntPercentage < bestBurntPercentage && totalBurntPercentage != 0) {
				bestBurntPercentage = totalBurntPercentage
				optimalDensity = density
			}
		}
		fmt.Printf("Tree density: %.2f, Average burnt percentage: %.2f%%\n", density, totalBurntPercentage/float64(trials))
	}
	fmt.Printf("optimal burnt percentage is %.2f%% for density of %.2f", bestBurntPercentage, optimalDensity)
}
