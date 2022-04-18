package main

import (
	"container/list"
	"fmt"
)

type treeNode struct {
	Source int
	Dest   int
}

/*
	複製垃圾收集，切尼算法。
	找到兩個節點u和之間的最短路徑v，路徑長度由邊的總數測量（優於深度優先搜索）。
	測試圖表的二分性。
	未加權圖的最小生成樹。
	網絡爬蟲。
	在圖的任何連接組件中查找節點。
	用於計算流網絡中最大流量的 Ford-Fulkerson 方法。
	二叉樹的序列化/反序列化與排序順序的序列化可以有效地重建樹。
*/

func main() {
	// create test tree
	tree := []treeNode{
		{1, 2}, {1, 3}, {1, 4}, {2, 5},
		{2, 6}, {5, 9}, {5, 10}, {4, 7},
		{4, 8}, {7, 11}, {7, 12},
	}

	// test total node
	n := 15
	graph := graph(tree, n)

	fmt.Printf("%v", graph)
	discovery := make([]bool, 15)

	for i := 0; i < n; i++ {
		fmt.Printf(" node : %v \n", i)
		if !discovery[i] {
			btw(graph, i, discovery)
		}
		fmt.Println("--------")
	}
}

//  node of tree connection
func graph(node []treeNode, num int) [][]int {
	dataTree := make([][]int, num)

	for i := range node {
		dataTree[node[i].Source] = append(dataTree[node[i].Source], node[i].Dest)
		dataTree[node[i].Dest] = append(dataTree[node[i].Dest], node[i].Source)
	}

	return dataTree
}

//  BTW
func btw(graph [][]int, node int, discovery []bool) {

	queue := list.New()
	queue.PushBack(node)

	// until queue empty
	for e := queue.Front(); e != nil; e = e.Next() {

		// if type of value is int
		if value, ok := e.Value.(int); ok {

			//check  the specify node's connections are  push to queue
			for _, v := range graph[value] {

				//if not in queue
				if !discovery[v] {
					discovery[v] = true
					queue.PushBack(v)
					fmt.Printf("%v, ", v)
				}
			}
			fmt.Printf("\n")
		}
	}
}
