package main

import (
	"algorithm/generact"
	"fmt"
	"log"
)

// worse : O(n^{2})
// avg   : O(nlog n)
// memory : O(log n)
// unstable

func main() {
	gen := generact.Int10
	log.Printf("%v", gen)
	sorted := quickSort(gen)

	log.Printf("%v", sorted)
}

// 快速排序
func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

// 遞歸分區 直到 left = right
func _quickSort(arr []int, left, right int) []int {
	if left < right {
		showDeital(left, right, arr)
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

// 分區 回傳分區節點 index
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		showPartition(pivot, i, index, arr)
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	if i != j {
		showSwap(i, j, arr)
	}
}

// 顯示詳細資料
func showDeital(left, right int, nums []int) {
	fmt.Printf("\x1B[44m %v \x1B[0m= > ", "d")
	for k := range nums {
		if k <= right && k >= left {
			fmt.Printf("\x1B[30m\x1B[43m %02d \x1B[0m", nums[k])
		} else {
			fmt.Printf(" %02d ", nums[k])
		}
	}

	fmt.Printf(" | i = %02d | j = %02d | ", left, right)

	fmt.Println("/t")
}

// 顯示詳細資料
func showPartition(pivot, i, index int, nums []int) {
	fmt.Printf(" %v ", "p = >")

	for k := range nums {
		if k < index {
			fmt.Printf("\x1B[30m\x1B[42m %02d \x1B[0m", nums[k])
		} else if k == i {
			fmt.Printf("\x1B[32m\x1B[44m %02d \x1B[0m", nums[k])
		} else if k < pivot {
			fmt.Printf("\x1B[30m\x1B[47m %02d \x1B[0m", nums[k])
		} else {
			fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
		}
	}

	fmt.Printf(" | pivot = %02d | i = %02d | index =  %02d ", pivot, i, index)

	fmt.Println("/t")
}

// 顯示詳細資料
func showSwap(left, right int, nums []int) {
	fmt.Printf(" %v ", "c = >")

	for k := range nums {
		if k == right || k == left {
			fmt.Printf("\x1B[31m\x1B[44m %02d \x1B[0m", nums[k])
		} else {
			fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
		}
	}

	fmt.Println(" | /t")
}
