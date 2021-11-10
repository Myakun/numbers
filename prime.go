package numbers

import (
	"fmt"
	"math"
)

// GetPrimeNumbers returns list of prime numbers from 2 to number
func GetPrimeNumbers(number int) ([]int, error) {
	if number <= 1 {
		return nil, fmt.Errorf("number need to be greater than 1, %d passed", number)
	}

	var primeNumbers []int
	primeNumbers = append(primeNumbers, 2)

	for i := 3; i <= number; i += 2 {
		isPrime := true
		for j := 3; j <= int(math.Sqrt(float64(i))); j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers, nil
}
