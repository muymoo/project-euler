package main

import (
	"fmt"
)

func main() {
	var sieve [110000]bool
	var primes []int

	for i := range sieve {
		sieve[i] = true
	}

	for i := 2; i < len(sieve); i++ {
		if sieve[i] {
			for j := i*i; j < len(sieve); j+=i {
				sieve[j] = false
				
			}
		}
	}

	for i, v := range sieve {
		if v {
			primes = append(primes, i)
		}
	}

	fmt.Println("size", len(primes))
	fmt.Println("10,001st", primes[10000])
}