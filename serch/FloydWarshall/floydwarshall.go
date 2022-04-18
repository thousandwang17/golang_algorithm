package main

import (
	"fmt"
	"math"
)

const I = math.MaxInt

func main() {

	graph := [][]int{
		{0, I, -2, I},
		{4, 0, 3, I},
		{I, I, 0, 2},
		{I, -1, I, 0},
	}

	floydWarshall(graph)
}

func floydWarshall(graph [][]int) {

	length := len(graph)
	path := make([][]int, length)
	cost := make([][]int, length)

	// init
	for i := 0; i < length; i++ {
		path[i] = make([]int, length)
		cost[i] = make([]int, length)

		for j := 0; j < length; j++ {
			cost[i][j] = graph[i][j]
			if i == j {
				path[i][j] = 0
			} else if cost[i][j] != I {
				path[i][j] = i
			} else {
				path[i][j] = -1
			}
		}
	}

	printCost(path, 0, 0, 0)

	for k := 0; k < length; k++ {
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if cost[i][j] > cost[i][k]+cost[k][j] && cost[i][k] != I && cost[k][j] != I {
					cost[i][j] = cost[i][k] + cost[k][j]
					path[i][j] = path[k][j]
					printCost(path, i, j, k)
				}

				if cost[i][i] < 0 {
					return
				}
			}
		}
	}

	printSolution(path)
}

func printSolution(path [][]int) {

	length := len(path)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i != j && path[i][j] != -1 {
				fmt.Printf("vertex %v to %v  |  %v ", i, j, i)
				printPath(path, i, j)
				fmt.Printf(" -> %v \n", j)
			}
		}
	}
}

func printCost(cost [][]int, v, u, k int) {
	len := len(cost)
	fmt.Printf(" i = %v j =  %v k = %v \n", v, u, k)
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if cost[i][j] == I {
				fmt.Printf(" I ,")
			} else {
				fmt.Printf("%02d ,", cost[i][j])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("--------\n")
}

func printPath(path [][]int, i, j int) {
	if path[i][j] == i {
		return
	}
	printPath(path, i, path[i][j])
	fmt.Printf(" -> %v ", path[i][j])
}
