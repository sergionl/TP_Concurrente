package main

import (
	"fmt"
	"sort"
	"time"
)

type node struct {
	value       int
	left, right *node
}

func insert(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	if value < root.value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}
	return root
}

func traverse(root *node, values []int) []int {
	if root != nil {
		values = traverse(root.left, values)
		values = append(values, root.value)
		values = traverse(root.right, values)
	}
	return values
}

func BinaryTreeSort(values []int) []int {
	var root *node
	for _, value := range values {
		root = insert(root, value)
	}
	return traverse(root, []int{})
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

	fmt.Printf("BinaryTreeSort (non-concurrent) took an average of %.6f seconds over %d tests with %d elements (trimmed mean)\n",
		trimmedMean, numTests, numElements)
}
