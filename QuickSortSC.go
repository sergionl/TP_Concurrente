package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func quickSort(values []int) {
	if len(values) < 2 {
		return
	}

	left, right := 0, len(values)-1
	pivot := rand.Int() % len(values)

	values[pivot], values[right] = values[right], values[pivot]

	for i := range values {
		if values[i] < values[right] {
			values[i], values[left] = values[left], values[i]
			left++
		}
	}

	values[left], values[right] = values[right], values[left]

	quickSort(values[:left])
	quickSort(values[left+1:])
}

func main() {
	numTests := 1000
	numElements := 1000000
	values := make([]int, numElements)

	for i := 0; i < numElements; i++ {
		values[i] = rand.Int()
	}

	var times []float64
	for i := 0; i < numTests; i++ {
		start := time.Now()
		quickSort(values)
		duration := time.Since(start)
		times = append(times, duration.Seconds())
	}

	sort.Float64s(times)

	var sum float64
	for i := 50; i < numTests-50; i++ {
		sum += times[i]
	}
	trimmedMean := sum / float64(numTests-100)

	fmt.Printf("Quick Sort took an average of %.6f seconds over %d tests with %d elements (trimmed mean)\n",
		trimmedMean, numTests, numElements)
}
