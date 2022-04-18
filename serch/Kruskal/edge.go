package main

type Edge struct {
	Src   int
	Dest  int
	Width int
}

type Edges []Edge

func (e Edges) Len() int { return len(e) }

func (e Edges) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func (e Edges) Less(i, j int) bool { return e[i].Width < e[j].Width }
