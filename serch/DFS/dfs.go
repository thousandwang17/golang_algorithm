package main

import (
	"fmt"
)

type treeNode struct {
	Source int
	Dest   int
}

/*
	前序 ： copy 副本
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

	fmt.Println("------------ post order ------------")
	for i := 0; i < n; i++ {
		fmt.Printf(" node : %v \n", i)
		if !discovery[i] {
			dtw(graph, i, discovery)
		}
	}

	fmt.Println("------------ pre order------------")

	discovery_pre := make([]bool, 15)

	for i := 0; i < n; i++ {
		if !discovery_pre[i] {
			dtw_pre(graph, i, discovery_pre)
		}
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

//  DTW - post order
// can delete tree
func dtw(graph [][]int, node int, discovery []bool) {

	discovery[node] = true

	for _, v := range graph[node] {
		//if not in queue
		if !discovery[v] {
			dtw(graph, v, discovery)
		}
	}
	fmt.Printf(" -> %v \n", node)

}

//  DTW - pre order
// can copy tree
func dtw_pre(graph [][]int, node int, discovery []bool) {

	discovery[node] = true
	fmt.Printf(" -> %v \n", node)

	for _, v := range graph[node] {
		//if not in queue
		if !discovery[v] {
			dtw_pre(graph, v, discovery)
		}
	}

}
