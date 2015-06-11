package main

import "fmt"

func main() {
	var numberToFactor, currentDivisor, largestDivisor int = 600851475143, 2, 0

	for numberToFactor > 1 {
		if numberToFactor % currentDivisor == 0 {
			numberToFactor = numberToFactor / currentDivisor
			largestDivisor = currentDivisor
			currentDivisor = 2
		} else {
			currentDivisor += 1
		}
	}
    fmt.Printf("%d\n", largestDivisor)
}