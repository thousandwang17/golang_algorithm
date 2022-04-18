package main

import (
	"algorithm/generact"
	"log"
	"math"
	"sort"
)

// binary search
// use case : when need return index that target equal value from sorted and unique slice

func main() {
	gen := generact.Int10
	sort.Ints(gen)

	log.Printf("0| %v", gen)
	log.Printf("1| %v", binary(2, gen))
	log.Printf("2| %v", binary(17, gen))
}

func binary(target int, nums []int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {

		mid := int(math.Floor(float64((end + start) / 2)))
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
