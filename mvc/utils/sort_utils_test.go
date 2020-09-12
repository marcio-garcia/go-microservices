package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortSuccess(t *testing.T) {
	// Initialization
	elements := []int{9, 8, 7, 6, 5}

	// Execution
	BubbleSort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

func TestBubbleSortWithNilInput(t *testing.T) {
	var elements []int = nil

	BubbleSort(elements)

	assert.Nil(t, elements)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	elements := getElements(5)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 4, elements[0])
	assert.EqualValues(t, 3, elements[1])
	assert.EqualValues(t, 2, elements[2])
	assert.EqualValues(t, 1, elements[3])
	assert.EqualValues(t, 0, elements[4])
}

// Benchmark - BubbleSort
func BenchmarkBubbleSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

// Benchmark - Go native sort
func BenchmarkNativeSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkNativeSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkNativeSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

// Benchmark - selecting one of the sort algorithms based on the number of elements
func BenchmarkSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
