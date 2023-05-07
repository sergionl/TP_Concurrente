package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type node struct {
	value       int
	left, right *node
}

func insert(root *node, value int, wg *sync.WaitGroup) *node {
	if root == nil {
		wg.Done()
		return &node{value: value}
	}
	if value < root.value {
		root.left = insert(root.left, value, wg)
	} else {
		root.right = insert(root.right, value, wg)
	}
	wg.Done()
	return root
}

func traverse(root *node, values chan<- int, wg *sync.WaitGroup) {
	if root != nil {
		traverse(root.left, values, wg)
		values <- root.value
		traverse(root.right, values, wg)
	}
	wg.Done()
}

func BinaryTreeSort(values []int) []int {
	var root *node
	var wgInsert sync.WaitGroup
	for _, value := range values {
		wgInsert.Add(1)
		go insert(root, value, &wgInsert)
	}
	wgInsert.Wait()

	valuesChan := make(chan int)
	var wgTraverse sync.WaitGroup
	wgTraverse.Add(1)
	go func() {
		traverse(root, valuesChan, &wgTraverse)
		close(valuesChan)
	}()

	sortedValues := make([]int, 0, len(values))
	for value := range valuesChan {
		sortedValues = append(sortedValues, value)
	}
	wgTraverse.Wait()

	return sortedValues
}

func main() {
	numTests := 1000
	numElements := 1000000
	values := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		values[i] = numElements - i
	}

	var times []float64
	for i := 0; i < numTests; i++ {
		start := time.Now()
		BinaryTreeSort(values)
		duration := time.Since(start)
		times = append(times, duration.Seconds())
	}

	sort.Float64s(times)
	var sum float64
	for i := 50; i < numTests-50; i++ {
		sum += times[i]
	}
	trimmedMean := sum / float64(numTests-100)

	fmt.Printf("BinaryTreeSort took an average of %.6f seconds over %d tests with %d elements (trimmed mean)\n",
		trimmedMean, numTests, numElements)
}
