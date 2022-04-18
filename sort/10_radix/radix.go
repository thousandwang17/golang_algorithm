package main

import (
	"algorithm/generact"
	"log"
)

// worse : O(n^2)
// avg   : O(k x n)
// stable

func main() {
	gen := generact.Int10

	log.Printf("%v", gen)
	result := radixSort(gen)

	log.Printf("%v", result)

}

func radixSort(nums []int) []int {
	max := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	len := 1
	for max_divide := 10; max/max_divide > 0; max_divide *= 10 {
		len++
	}

	base := 1
	divide := 10
	radixSlice := make([][]int, 10)

	for i := 0; i < len; i++ {

		for j := range nums {
			index := (nums[j] % divide) / base
			radixSlice[index] = append(radixSlice[index], nums[j])
		}

		nums_index := 0
		for r := range radixSlice {
			for rr := range radixSlice[r] {
				nums[nums_index] = radixSlice[r][rr]
				nums_index++
			}
			radixSlice[r] = radixSlice[r][:0]
		}

		divide *= 10
		base *= 10
	}

	return nums
}
