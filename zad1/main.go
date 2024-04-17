package main

//Kajetan Wiśniewski

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
)

func main() {

	var generatedNickname string
	var err error
	var nicknameInASCII [6]int

	for {
		fmt.Println("wpisz swoje imie i nazwisko rozdzielone spacja aby wygenerowac nickname:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		generatedNickname, err = generateNickname(input)
		if err == nil {
			nicknameInASCII = nicknameToASCII(generatedNickname)
			break
		}
		fmt.Println("Error:", err)
	}

	n := big.NewInt(1)
	strongNumber := 1

	for {
		factorialResult := factorial(n)
		if containsAllDigits(factorialResult, nicknameInASCII) {
			break
		}
		n.Add(n, big.NewInt(1))
		strongNumber++
	}

	callCount := make(map[int]int)
	for i := 30; i >= 1; i-- {
		fibonacciWithCalls(i, callCount)
	}

	minDiff := strongNumber
	weakNumber := 0
	weakNumber2 := 0

	for i := 0; i <= 30; i++ {
		diff := abs(callCount[i]+1 - strongNumber)
		if diff < minDiff {
			minDiff = diff
			weakNumber = i
			weakNumber2 = callCount[i] + 1
		}
	}
	
	fmt.Printf("Silna liczba dla %s to %d a Slaba liczba to %d [dla %d wywolan]\n", generatedNickname, strongNumber, weakNumber, weakNumber2)
	time.Sleep(5 * time.Second) //to tylko po to aby przy uzywaniu pliku .exe wyswietlic wynik bez
	//dodatkowego nieskonczonego for loopa itd
}
