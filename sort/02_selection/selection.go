package main

import (
	"algorithm/generact"
	"fmt"
)

// worse : O(n^{2})
// avg   : O(n^{2})
// unstable
// 冒泡眼算法 bubble sort
func main() {
	// nums := generact.IntSlice(10)
	nums := generact.Int10
	count := int(1)

	// 每次 i 遍例時 將最大值往前排序 , 直到最後為止
	for i := 0; i < len(nums)-1; i++ {
		//
		min_index := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
			selectionDeital(count, i, j, min_index, nums, false)
			count++
		}
		selectionDeital(count, i, min_index, min_index, nums, true)
		nums[i], nums[min_index] = nums[min_index], nums[i]
	}
}

// 顯示詳細資料
func selectionDeital(count, i, j, min_index int, nums []int, change bool) {
	fmt.Printf("%03d => ", count)
	for k := range nums {
		if change && k == i {
			fmt.Printf("\x1B[43m\x1B[31m %02d \x1B[0m", nums[k])
		} else if k == j {
			// 當前比較
			if j == min_index {
				if change {
					fmt.Printf("\x1B[43m\x1B[31m %02d \x1B[0m", nums[k])
				} else {
					fmt.Printf("\x1B[31m\x1B[47m %02d \x1B[0m", nums[k])
				}
			} else {
				// 當前比較欄位
				fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
			}
		} else if k == min_index {
			fmt.Printf("\x1B[31m\x1B[47m %02d \x1B[0m", nums[k])
		} else if k < i {
			// 不用在比較 紫底顯示
			fmt.Printf("\x1B[44m %02d \x1B[0m", nums[k])
		} else if k < j && !change {
			fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
		} else {
			// 非當前比較
			fmt.Printf(" %02d ", nums[k])
		}
	}

	if change {
		fmt.Printf(" | changed                   ")
	} else {
		fmt.Printf(" | i = %02d | j = %02d | min = %02d", i, j, nums[min_index])
	}

	fmt.Println("  /t")
}

// 運算複雜度 n2
// 空間服札度 n
