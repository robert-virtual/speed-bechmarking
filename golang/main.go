package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Expected two arguments example: ./main 2 100")
		os.Exit(1)
	}
	start, _ := strconv.ParseFloat(os.Args[1], 64)
	end, _ := strconv.ParseFloat(os.Args[2], 64)
	primes := []float64{}
	for n := start; n <= end; n++ {
		root := math.Sqrt(n)
		isPrime := true
		for i := 2.0; i <= root; i++ {

			if math.Mod(n, i) == 0 {
				isPrime = false
			}

		}
		if isPrime {
			primes = append(primes, n)
		}
	}
	fmt.Printf("%v primes found from %v to %v \n", len(primes), start, end)
	for i := 0; i < len(primes); i++ {
		fmt.Printf("%v) %v \n", i, primes[i])
	}

}
