package sort

import (
	"math/rand"
	"time"
)

func BubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

func SelectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		var minIndex = i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}

		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

func InsertionSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := 1; i < len(ar); i++ {
		for j := i; j > 0 && ar[j-1] > ar[j]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}
}
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func QuickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}
	rand.Seed(time.Now().UnixNano())

	pivotIndex := rand.Intn(len(ar))
	pivot := ar[pivotIndex]
	ar[pivotIndex] = ar[len(ar)-1]
	ar[len(ar)-1] = pivot
	left := -1

	for right := range ar {
		if ar[right] < pivot {
			left++
			swap(ar, left, right)
		}
	}
	swap(ar, left+1, len(ar)-1)

	QuickSort(ar[:left+1])
	QuickSort(ar[left+2:])
}

func swap(ar []int, i, j int) {
	ar[i], ar[j] = ar[j], ar[i]
}
