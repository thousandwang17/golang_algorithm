package main

import (
	"algorithm/generact"
	"fmt"
)

// worse : O(n^{2})
// avg   : O(n^{2})
// unstable

func main() {
	// num := generact.IntSlice(10)
	num := generact.Int10
	count := 0

	for i := range num {
		preIndex := i - 1
		current := num[i]

		// if current num(i) > pre num(i-1) ,change eachother
		// until num's index equl zero or pre value small than current
		for preIndex >= 0 && num[preIndex] > current {
			showDeital(count, i, preIndex, current, num, false, true)
			num[preIndex+1] = num[preIndex] //
			showDeital(count, i, preIndex, current, num, true, true)
			preIndex -= 1
			count++
		}

		// at the end set current value
		num[preIndex+1] = current
		showDeital(count, i, preIndex, current, num, false, false)
		count++

	}
}

// 顯示詳細資料
func showDeital(count, i, j, current int, nums []int, change, inWhile bool) {
	fmt.Printf("%03d => ", count)
	for k := range nums {
		if k == j {
			if inWhile && change {
				fmt.Printf("\x1B[31m\x1B[44m %02d \x1B[0m", current)
			} else if change {
				fmt.Printf("\x1B[32m\x1B[44m %02d \x1B[0m", nums[k])
			} else {
				fmt.Printf("\x1B[32m\x1B[44m %02d \x1B[0m", nums[k])
			}
		} else if k == j+1 {
			if inWhile && !change {
				fmt.Printf("\x1B[32m\x1B[44m %02d \x1B[0m", current)
			} else if change {
				fmt.Printf("\x1B[31m\x1B[44m %02d \x1B[0m", nums[k])
			} else {
				fmt.Printf("\x1B[32m\x1B[44m %02d \x1B[0m", nums[k])
			}
		} else if k <= i {
			fmt.Printf("\x1B[44m %02d \x1B[0m", nums[k])
		} else {
			// 非當前比較
			fmt.Printf(" %02d ", nums[k])
		}
	}
	if change {
		fmt.Printf(" | changed           ")
	} else {
		fmt.Printf(" | i = %02d | j = %02d | ", i, j)
	}
	fmt.Println("/t")
}
