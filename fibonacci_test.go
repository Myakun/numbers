package numbers_test

import (
	"fmt"
	"github.com/myakun/numbers"
	"testing"
)

type testItem struct {
	expected uint
	number   uint
}

func ExampleFib() {
	fmt.Println(numbers.Fib(7))
	// Output: 13
}

func TestFibResult(t *testing.T) {
	tests := []testItem{
		{expected: 0, number: 0},
		{expected: 1, number: 1},
		{expected: 1, number: 2},
		{expected: 2, number: 3},
		{expected: 3, number: 4},
		{expected: 5, number: 5},
		{expected: 8, number: 6},
		{expected: 13, number: 7},
		{expected: 21, number: 8},
		{expected: 34, number: 9},
	}

	for _, test := range tests {
		if res := numbers.Fib(test.number); res != test.expected {
			t.Errorf("Incorrect number for input %d: expected %d, got %d", test.number, test.expected, res)
		}
	}
}
