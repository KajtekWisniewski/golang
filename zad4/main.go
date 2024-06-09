package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// inicjalizacja kolektora colly
	c := colly.NewCollector()

	var allRows [][]string
	tableCount := 0

	// funkcja ktora wykonujemy przy znalezieni kazdego elementu HTML <table> <tbody>
	c.OnHTML("table.wikitable > tbody", func(h *colly.HTMLElement) {
		tableCount++
		if tableCount != 2 {
			return
		}

		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var row []string
			el.ForEach("th, td", func(_ int, col *colly.HTMLElement) {
				row = append(row, strings.TrimSpace(col.Text))
			})
			if len(row) > 0 {
				allRows = append(allRows, row)
			}
		})
	})

	// url strony z ktorej pobieramy dane
	c.Visit("https://en.wikipedia.org/wiki/List_of_European_Cup_and_UEFA_Champions_League_finals")

	// sprawdzenie czy dane zostaly pobrane
	if len(allRows) == 0 {
		fmt.Println("No data found")
		return
	}

	//zapis do csv
	file, err := os.Create("zwyciezcy_uefa_cl.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Zapis danych
	for _, row := range allRows {
		writer.Write(row)
	}

	fmt.Println("dane zostaly zapisane do zwyciezcy_uefa_cl.csv")
}