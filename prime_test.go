package numbers_test

import (
	"fmt"
	"github.com/myakun/numbers"
	"testing"
)

func ExampleGetPrimeNumbers() {
	fmt.Println(numbers.GetPrimeNumbers(8))
	// Output: [2 3 5 7 11 13 17 19 23]
}

func TestGetPrimeNumbersCorrectResult(t *testing.T) {
	expected := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23}

	primeNumbers, err := numbers.GetPrimeNumbers(expected[len(expected)-1])

	if nil != err {
		t.Fatalf("Unexpected error: %s", err)
	}

	if len(primeNumbers) != len(expected) {
		t.Fatalf("Incorrect result slice length. Expected %d, got %d", len(expected), len(primeNumbers))
	}

	for i, v := range expected {
		if v != primeNumbers[i] {
			t.Errorf("Incorrect prime number at position %d: expected %d, got %d", i, v, primeNumbers[i])
		}
	}
}

func TestGetPrimeNumbersIncorrectInput(t *testing.T) {
	_, err := numbers.GetPrimeNumbers(1)

	if nil == err {
		t.Errorf("Expected error, got nil")
	}
}
