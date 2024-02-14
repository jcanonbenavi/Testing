package main

import (
	"fmt"
	"unit_testing_primes/internal"
)

func main() {
	prime, err := internal.PrimeNumber(31)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prime)

}
