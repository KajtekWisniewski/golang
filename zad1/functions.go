package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func generateNickname(input string) (string, error) {

    parts := strings.Split(input, " ")
    if len(parts) != 2 {
        return "", fmt.Errorf("input powinien zawierac imie i nazwisko rozdzielone spacja")
    }
    name := strings.ToLower(parts[0])
    surname := strings.ToLower(parts[1])
    var nickname string
    if len(name) >= 3 && len(surname) >= 3 {
        nickname = name[:3] + surname[:3]
    } else {
        return "", fmt.Errorf("imie i nazwisko, oba, powinny miec przynajmniej 3 znaki dlugosci")
    }

    return nickname, nil
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func nicknameToASCII(nickname string) [6]int {
    var asciiArray [6]int

    for i := 0; i < 6; i++ {
        if i < len(nickname) {
            asciiArray[i] = int(nickname[i])
        } else {
            asciiArray[i] = 0
        }
    }

    return asciiArray
}

func containsAllDigits(factorialResult *big.Int, digits [6]int) bool {
	for _, digit := range digits {
		if !containsSubNumber(factorialResult, digit) {
			return false
		}
	}
	return true
}

func containsSubNumber(bigNum *big.Int, subNum int) bool {
    bigNumStr := bigNum.String()
    subNumStr := strconv.Itoa(subNum)
    return strings.Contains(bigNumStr, subNumStr)
}

func fibonacciWithCalls(n int, callCount map[int]int) int {
	if n < 0 {
		callCount[n] = 1
		return 1
	}

	callCount[n-2]++

	return fibonacciWithCalls(n-1, callCount) + fibonacciWithCalls(n-2, callCount)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}