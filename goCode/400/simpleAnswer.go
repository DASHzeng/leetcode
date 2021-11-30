package main

import "fmt"

func findNthDigit(n int) int {
	source := 9
	i := 1
	for n > i*source {
		n -= i * source
		source *= 10
		i += 1
	}
	source /= 9
	for n > i {
		source += 1
		n -= i
	}
	for n < i {
		source /= 10
		n += 1
	}
	return source % 10
}

func main() {
	fmt.Print(findNthDigit(100))
}
