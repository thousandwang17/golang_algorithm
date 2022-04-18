package main

import (
	"algorithm/generact"
	"fmt"
)

// worse :  O(n^{2})
// avg   : O(n^{2})
// stable
// 冒泡眼算法 bubble sort
func main() {
	nums := generact.IntSlice(10)
	count := int(1)

	// 每次 i 遍例時 將最大值往後排序 , 直到上筆排序最大值之前為止
	for i := 0; i < len(nums); i++ {
		// 每次 從頭比對相鄰兩位值 如果 前值大於後 則交換兩邊的值 , 直到上次排序結果之前
		for j := 0; j < len(nums)-1-i; j++ {
			showDeital(count, i, j, nums, false)
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				showDeital(count, i, j, nums, true)
			}
			count++
		}
	}
}

// 顯示詳細資料
func showDeital(count, i, j int, nums []int, change bool) {
	fmt.Printf("%03d => ", count)
	for k := range nums {
		if k == j || k == j+1 {
			// 當前比較
			if change {
				// 如果有交換 紅字顯示
				fmt.Printf("\x1B[31m %02d \x1B[0m", nums[k])
			} else {
				// 當前比較欄位
				fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
			}
		} else if k > len(nums)-1-i {
			// 不用在比較 紫底顯示
			fmt.Printf("\x1B[44m %02d \x1B[0m", nums[k])
		} else {
			// 非當前比較
			fmt.Printf(" %02d ", nums[k])
		}
	}
	fmt.Printf(" | i = %d | j = %d | ", i, j)
	fmt.Println("/t")
}
