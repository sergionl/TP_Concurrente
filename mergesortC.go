package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func mergeSortConcurrent(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftChan := make(chan []int)
	rightChan := make(chan []int)

	// Divide el arreglo en subarreglos y los ordena de forma concurrente
	go func() {
		leftChan <- mergeSortConcurrent(arr[:mid])
	}()

	go func() {
		rightChan <- mergeSortConcurrent(arr[mid:])
	}()

	// Combina los subarreglos ordenados
	left, right := <-leftChan, <-rightChan
	close(leftChan)
	close(rightChan)

	return mergeConcurrent(left, right)
}

func mergeConcurrent(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numTrials := 1000
	arrSize := 1000000

	// Genera el arreglo de prueba una sola vez
	arr := rand.Perm(arrSize)

	// Almacena los tiempos de ejecución
	times := make([]time.Duration, numTrials)

	for i := 0; i < numTrials; i++ {
		// Crea una copia del arreglo de prueba
		arrCopy := make([]int, arrSize)
		copy(arrCopy, arr)

		// Ejecuta el ordenamiento y registra el tiempo de ejecución
		startTime := time.Now()
		_ = mergeSortConcurrent(arrCopy)
		times[i] = time.Since(startTime)
	}

	// Ordena los tiempos para calcular la media recortada
	sort.Slice(times, func(i, j int) bool {
		return times[i] < times[j]
	})

	// Calcula la media recortada omitiendo los 50 menores y los 50 mayores tiempos
	var total time.Duration
	for i := 50; i < numTrials-50; i++ {
		total += times[i]
	}
	average := total / time.Duration(numTrials-100)

	fmt.Println("Average execution time:", average)
}

//En esta implementación, la función mergeSortConcurrent() divide el arreglo en dos subarreglos y utiliza goroutines para ordenarlos de forma concurrente. Luego, utiliza la función mergeConcurrent() para combinar los subarreglos ordenados en un solo arreglo ordenado.
//
//La función mergeConcurrent() es similar a la versión sin concurrencia, pero no utiliza un bucle while para combinar los subarreglos, ya que esto no es adecuado para la concurrencia. En su lugar, utiliza un bucle for que se ejecuta mientras hay elementos en ambos subarreglos.
//
//Finalmente, la función main() genera un arreglo de 10 números aleatorios sin repetición utilizando rand.Perm() y utiliza mergeSortConcurrent() para ordenarlo.
