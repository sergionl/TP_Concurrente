package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
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
		_ = mergeSort(arrCopy)
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

//Esta implementación de Merge Sort en Go utiliza recursión para dividir el arreglo en subarreglos más pequeños hasta que cada subarreglo contenga un solo elemento. Luego, utiliza la función merge para combinar los subarreglos ordenados en un solo arreglo ordenado.
//
//La función merge toma dos subarreglos ordenados como entrada y devuelve un arreglo ordenado que contiene todos los elementos de los dos subarreglos. La función utiliza un bucle while para comparar los primeros elementos de cada subarreglo y agregar el menor de los dos al resultado final.
//
//Finalmente, la función mergeSort devuelve el arreglo ordenado completo llamando a la función merge con los subarreglos izquierdo y derecho.

//En este ejemplo, la función main utiliza la función rand.Perm() de la biblioteca estándar de Go para generar un arreglo de 10 números aleatorios sin repetición. Luego, imprime el arreglo sin ordenar y lo ordena utilizando la función mergeSort que definimos anteriormente. Finalmente, imprime el arreglo ordenado.
