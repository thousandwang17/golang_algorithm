package main

import (
	"algorithm/generact"
	"fmt"
	"log"
)

// worse : O(nlog n)
// avg   : O(nlog n)
// unstable

func main() {
	gen := generact.Int10
	log.Printf("%v", gen)
	sorted := heapSort(gen)

	log.Printf("%v", sorted)
}

func heapSort(arr []int) []int {
	arrLen := len(arr)

	// init
	buildMaxHeap(arr, arrLen)

	fmt.Println("--------------- End buildMaxHeap ---------------")

	//  at this while condiction  the first value always is biggest unsort number
	//  swap first value and arrLen index value to be sort number
	//  until the index is zero
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

// init binary tree  that  move forward the bigger number at all tree
// start at the last node  , while until root node
func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

// recursive that move forward the max number
func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1  // left sub tree index
	right := 2*i + 2 // right sub tree index
	largest := i

	// if left index  bigger  than  arrLen index then pass
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}

	// if right index  bigger  than  arrLen index then pass
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}

	showHeapify(left, right, arrLen, i, arr)

	// if largest index equel i index that mean under the node all
	// order by number size
	// otherwise  recursive all sub tree sequentially
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	showSwap(i, j, arr)
}

// 顯示詳細資料
func showHeapify(left, right, arrLen, largest int, nums []int) {
	fmt.Printf(" %v ", "p = >")

	for k := range nums {
		if k >= arrLen {
			fmt.Printf(" %02d ", nums[k])
		} else if k == largest {
			fmt.Printf("\x1B[30m\x1B[43m %02d \x1B[0m", nums[k])
		} else if k == left || k == right {
			fmt.Printf("\x1B[30m\x1B[42m %02d \x1B[0m", nums[k])
		} else {
			fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
		}
	}

	fmt.Printf(" | left = %02d | right = %02d | arrLen  = %02d | largest =  %02d ", left, right, arrLen, largest)

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
