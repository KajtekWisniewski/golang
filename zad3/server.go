package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

type SharkAttack struct {
	Date     string `json:"date"`
	Country  string `json:"country"`
	Area     string `json:"area"`
	Location string `json:"location"`
	Activity string `json:"activity"`
	Injury   string `json:"injury"`
}

var (
	data     []SharkAttack
	dataLock sync.Mutex
)

func loadData() {
	file, err := os.Open("global-shark-attack.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var allData []SharkAttack
	json.Unmarshal(byteValue, &allData)

	rand.Shuffle(len(allData), func(i, j int) { allData[i], allData[j] = allData[j], allData[i] })
	if len(allData) > 10 {
		allData = allData[:10]
	}

	dataLock.Lock()
	data = allData
	dataLock.Unlock()
}

func main() {
	const port = ":8080"
	loadData()
	http.HandleFunc("/posts", postsHandler)
	fmt.Println("serwer dziala na porcie", port)

	log.Fatal(http.ListenAndServe(port, nil))

}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		downloadHandler(w, r)
	case http.MethodPost:
		publishHandler(w, r)
	case http.MethodDelete:
		deleteHandler(w, r)
	default:
		http.Error(w, "Niewlasciwy request", http.StatusMethodNotAllowed)
	}
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Niewlasciwy request", http.StatusMethodNotAllowed)
		return
	}

	var attack SharkAttack
	if err := json.NewDecoder(r.Body).Decode(&attack); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dataLock.Lock()
	data = append(data, attack)
	dataLock.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "post zostal opublikowany")
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Niewlasciwy request", http.StatusMethodNotAllowed)
		return
	}

	dataLock.Lock()
	defer dataLock.Unlock()

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Niewlasciwy request", http.StatusMethodNotAllowed)
		return
	}

	var target SharkAttack
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dataLock.Lock()
	defer dataLock.Unlock()

	for i, attack := range data {
		if attack.Date == target.Date && attack.Country == target.Country && attack.Area == target.Area && attack.Location == target.Location {
			data = append(data[:i], data[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "post zostal usuniety")
			return
		}
	}

	http.Error(w, "nie znaleziono postu", http.StatusNotFound)
}
