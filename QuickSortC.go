package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func quickSort(arr []int, wg *sync.WaitGroup) {
	if len(arr) <= 1 {
		wg.Done()
		return
	}

	pivot := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		} else {
			arr[right], arr[i] = arr[i], arr[right]
			right--
		}
	}

	wgLeft := sync.WaitGroup{}
	wgLeft.Add(1)
	wgRight := sync.WaitGroup{}
	wgRight.Add(1)

	go quickSort(arr[:left], &wgLeft)
	go quickSort(arr[left+1:], &wgRight)

	wgLeft.Wait()
	wgRight.Wait()
	wg.Done()
}

func quickSortConcurrent(arr []int) {
	var wg sync.WaitGroup
	wg.Add(1)
	quickSort(arr, &wg)
	wg.Wait()
}

func main() {
	numTests := 1000
	numElements := 1000000
	values := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		values[i] = rand.Intn(numElements)
	}

	var times []float64
	for i := 0; i < numTests; i++ {
		start := time.Now()
		quickSortConcurrent(values)
		duration := time.Since(start)
		times = append(times, duration.Seconds())
	}

	sort.Float64s(times)
	var sum float64
	for i := 50; i < numTests-50; i++ {
		sum += times[i]
	}
	trimmedMean := sum / float64(numTests-100)

	fmt.Printf("Quick sort took an average of %.6f seconds over %d tests with %d elements (trimmed mean)\n",
		trimmedMean, numTests, numElements)
}
