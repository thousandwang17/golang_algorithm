package main

import (
	"algorithm/generact"
	"log"
	"math"
)

// worse :  O(n^2)
// avg   :  O( n)
// menory  : O(m)
// stable

func main() {
	gen := generact.Int10

	log.Printf("%v", gen)
	result := bucketSort(gen, 10)
	log.Printf("%v", result)

}

func bucketSort(nums []int, buckerSize int) []int {

	if len(nums) == 0 {
		return nums
	}

	max, min := 0, 0

	// find max and min value of nums
	for i := range nums {
		if min > nums[i] {
			min = nums[i]
		} else if max < nums[i] {
			max = nums[i]
		}
	}

	if buckerSize <= 0 {
		buckerSize = 5
	}

	// init bucket slice
	bucketCount := int(math.Floor(float64((max-min)/buckerSize))) + 1
	bucketSlice := make([][]int, bucketCount)

	// copy to bucket slice
	for i := range nums {
		index := int(math.Floor(float64((nums[i] - min) / buckerSize)))
		bucketSlice[index] = append(bucketSlice[index], nums[i])
	}

	index := 0
	// take from each bucket and sort
	for i := range bucketSlice {
		sorted := insertSort(bucketSlice[i])
		// replace sorted slice to nums
		for si := range sorted {
			nums[index] = sorted[si]
			index++
		}
	}

	return nums
}

func insertSort(nums []int) []int {

	for i := range nums {
		current := nums[i]
		preIndex := i - 1
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = current
	}

	return nums
}
