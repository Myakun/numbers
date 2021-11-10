package numbers

// InsertionSort returns sorted slice
func InsertionSort(numbers []int64) (sortedNumbers []int64) {
	sortedNumbers = make([]int64, len(numbers))
	copy(sortedNumbers, numbers)

	for i := 1; i < len(sortedNumbers); i++ {
		for j := i; j > 0 && sortedNumbers[j-1] > sortedNumbers[j]; j-- {
			buff := sortedNumbers[j]
			sortedNumbers[j] = sortedNumbers[j-1]
			sortedNumbers[j-1] = buff
		}
	}

	return
}
