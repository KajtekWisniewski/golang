package main

import (
	"fmt"
	"math/big"
)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func containsAllDigits(factorialResult *big.Int, digits []int) bool {
	for _, digit := range digits {
		if !containsDigit(factorialResult, digit) {
			return false
		}
	}
	return true
}

func containsDigit(number *big.Int, digit int) bool {
	for _, ch := range number.String() {
		if int(ch-'0') == digit {
			return true
		}
	}
	return false
}

func main() {
	n := big.NewInt(1)
	comparisonValue := big.NewInt(500)
	nickname := []int{112, 105, 111, 97, 114, 108}
	for {
		fmt.Printf("now calculating for n: %d\n", n)
		factorialResult := factorial(n)
		fmt.Println(containsAllDigits(factorialResult, nickname))
		if containsAllDigits(factorialResult, nickname) {
			fmt.Printf("Strong Number: %d\n", n)
			break
		}
		if n.Cmp(comparisonValue) == 1 {
			break
		}
		n.Add(n, big.NewInt(1))

	}
}
