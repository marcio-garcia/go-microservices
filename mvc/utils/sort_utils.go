package utils

import "sort"

// BubbleSort sorts a slice of integers using the bubblesort algorithm
func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

// Sort - receives a slice of integers and returns it in ascending order
func Sort(elements []int) {
	if len(elements) < 5000 {
		BubbleSort(elements)
		return
	}
	sort.Ints(elements)
}
