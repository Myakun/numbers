package numbers

var cache map[uint]uint

// Fib returns Fibonacci number
func Fib(number uint) uint {
	if nil == cache {
		cache = make(map[uint]uint)
	}

	if 0 == number {
		return 0
	} else if 1 == number {
		return 1
	}

	_, cacheExists := cache[number]
	if !cacheExists {
		cache[number] = Fib(number-1) + Fib(number-2)
	}

	return cache[number]
}
