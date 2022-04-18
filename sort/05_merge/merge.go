package main

import (
	"algorithm/generact"
	"log"
)

// worse : O(nlog n)
// avg   : O(nlog n)
// memory : O(n)+O(log n)
// stable

func main() {
	gen := generact.Int10
	log.Printf("%v", gen)
	sorted := mergeSort(gen)

	log.Printf("%v", sorted)
}

func mergeSort(num []int) []int {
	numlen := len(num)
	if numlen < 2 {
		return num
	}
	middle := numlen / 2
	left := num[:middle]
	right := num[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	total := len(left) + len(right)
	result := make([]int, 0, total)

	// var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
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
