package main

type UnorderedMap struct {
	parent []int
	rank   []int
}

func (u *UnorderedMap) MakeSet(number int) {
	u.parent = make([]int, number)
	u.rank = make([]int, number)

	for i := 0; i < number; i++ {
		u.parent[i] = i
		u.rank[i] = 0
	}
}

func (u *UnorderedMap) Find(k int) int {
	if u.parent[k] != k {
		u.parent[k] = u.Find(u.parent[k])
	}
	return u.parent[k]
}

func (u *UnorderedMap) Union(a, b int) {
	x := u.Find(a)
	y := u.Find(b)

	if x == y {
		return
	}

	if u.rank[x] > u.rank[y] {
		u.parent[y] = x
	} else if u.rank[x] < u.rank[y] {
		u.parent[x] = y
	} else {
		u.parent[x] = y
		u.rank[y]++
	}
}
