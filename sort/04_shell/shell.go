package main

import (
	"algorithm/generact"
	"fmt"
	"log"
)

// worse : O(n^{2})
// avg   : O(nlog ^{2}n)
// unstable

func main() {
	// gen := generact.IntSlice(100)
	shellSort(generact.Int10)
}

func shellSort(arr []int) []int {
	count := 1
	length := len(arr)
	gap := 1

	// gap 最優化
	for gap < length/3 {
		gap = gap*3 + 1
	}

	fmt.Printf("int => %v ", arr)
	fmt.Println(" /t ")

	// gap > 1 時 會先宏觀比較換位 , 減少微觀(gap = 1 )換位次數
	// gap = 1 時 會從頭遍例一次插入排序
	for gap > 0 {

		// gap 遍例到最後
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap

			pre := count

			for j >= 0 && arr[j] > temp {
				showDeital(count, i, j, gap, temp, arr, false)
				arr[j+gap] = arr[j]
				showDeital(count, i, j, gap, temp, arr, true)
				j -= gap
				count++
			}

			// if not change than call showDeital()
			if pre == count {
				showDeital(count, i, j, gap, temp, arr, false)
			}

			arr[j+gap] = temp
			count++
		}
		gap = gap / 3
		count++
	}
	log.Printf("%v", count)
	return arr
}

// 顯示詳細資料
func showDeital(count, i, j, gap, temp int, nums []int, change bool) {
	fmt.Printf("%03d => ", count)
	for k := range nums {
		if change {
			if j >= 0 {
				if k == j+gap {
					fmt.Printf("\x1B[31m\x1B[44m %02d \x1B[0m", nums[k])
				} else if k == j {
					fmt.Printf("\x1B[31m\x1B[44m %02d \x1B[0m", temp)
				} else {
					fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
				}
			}
		} else {
			if j >= 0 {
				if k == j+gap {
					fmt.Printf("\x1B[32m\x1B[44m %02d \x1B[0m", temp)
				} else if k == j {
					fmt.Printf("\x1B[32m\x1B[44m %02d \x1B[0m", nums[k])
				} else {
					fmt.Printf("\x1B[32m %02d \x1B[0m", nums[k])
				}
			}
		}
	}
	if change {
		fmt.Printf(" | changed                       ")
	} else {
		fmt.Printf(" | i = %02d | j = %02d |  gap = %02d | ", i, j, gap)
	}
	fmt.Println("/t")
}
