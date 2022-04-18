package main

import (
	"algorithm/generact"
	"log"
)

// worse : O(n+m)
// avg   : O(n+m)
// menory  : O(n+m)
// stable
func main() {
	gen := generact.Int10

	countSort(gen)
	log.Printf("%v", gen)
}

func countSort(nums []int) []int {

	// find max value
	maxValue := 0
	for i := range nums {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}

	// and create maxValue of len slice
	arr := make([]int, maxValue+1)
	arrIndex := 0

	// count value to index
	for i := range nums {
		arr[nums[i]] += 1
	}

	// print sorted arr until value in arr[i] equal to zero
	for i := range arr {
		for arr[i] > 0 {
			nums[arrIndex] = i
			arr[i]--
			arrIndex++
		}
	}
	return nums
}
