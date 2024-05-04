package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Empty = 0
	Tree  = 1
	Burnt = 2
)

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
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				spreadFire(forest, x+dx, y+dy)
			}
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
	rand.Seed(time.Now().UnixNano())
	size := 10
	treeProbability := 0.55

	forest := initForest(size, treeProbability)
	printForest(forest)
	startX, startY := rand.Intn(size), rand.Intn(size)
	spreadFire(forest, startX, startY)
	burntTrees := countBurntTrees(forest)
	emptySpaces := countEmptyLots(forest)

	fmt.Println("Forest after fire:")
	printForest(forest)
	fmt.Printf("Percentage of burnt forest: %.2f%%\n", 100*float64(burntTrees)/(float64(size*size)-float64(emptySpaces)))
	fmt.Println(burntTrees, emptySpaces)
}
