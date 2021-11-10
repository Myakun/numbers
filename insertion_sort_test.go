package numbers_test

import (
	"fmt"
	"github.com/myakun/numbers"
	"testing"
)

func ExampleInsertionSort() {
	data := []int64{11, 4, 3, 6, 9, 2}
	fmt.Println(numbers.InsertionSort(data))
	// Output: [2, 3, 4, 6, 9, 11]
}

func TestInsertionSort(t *testing.T) {
	input := []int64{11, 4, 3, 6, 9, 2}
	expected := []int64{2, 3, 4, 6, 9, 11}

	sorted := numbers.InsertionSort(input)

	if len(input) != len(sorted) {
		t.Fatalf("Incorrect result slice length. Expected %d, got %d", len(input), len(sorted))
	}

	for i, v := range expected {
		if v != sorted[i] {
			t.Errorf("Incorrect value at position %d: expected %d, got %d", i, v, sorted[i])
		}
	}
}
