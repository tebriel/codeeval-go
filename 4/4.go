package main

import "fmt"
import "math/big"

func main() {
    primes_found := 0
    primes_sum := 0
    for i := 1; ; i++ {
        if big.NewInt(int64(i)).ProbablyPrime(20) {
            primes_sum += i
            primes_found += 1
        }

        if primes_found == 1000 {
            break
        }
    }
    fmt.Println(primes_sum)
}
