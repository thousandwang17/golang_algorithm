package main

import (
	"fmt"
	"sort"
)

func main() {
	// create test tree
	tree := []Edge{
		{0, 1, 7}, {1, 2, 8}, {0, 3, 5}, {1, 3, 9}, {1, 4, 7}, {2, 4, 5},
		{3, 4, 15}, {3, 5, 6}, {4, 5, 8}, {4, 6, 9}, {5, 6, 11},
	}

	// test total node
	n := 7

	sort.Sort(Edges(tree))

	kruskals(tree, n)
}

func kruskals(tree []Edge, n int) {
	MST := make([]Edge, 0, n-1)
	treeLen := len(tree)

	subset := new(UnorderedMap)
	subset.MakeSet(n)

	fmt.Printf(" init => %+v \n", subset)

	index, i := 0, 0
	for index < n-1 && i < treeLen {

		e := tree[i]
		x := subset.Find(e.Src)
		y := subset.Find(e.Dest)

		if x != y {
			fmt.Printf("%v => ", e)
			MST = append(MST, e)
			subset.Union(x, y)
			fmt.Printf("%+v \n", subset)
			index++
		}
		i++
	}

	fmt.Printf("MST = %v ", MST)

}
