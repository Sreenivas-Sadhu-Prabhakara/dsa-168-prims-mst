package main

import (
	"container/heap"
	"fmt"
)

// Edge is an undirected weighted edge.
type Edge struct {
	U, V, W int
}

type pqItem struct{ w, v int }

type pq []pqItem

func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return p[i].w < p[j].w }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x any)        { *p = append(*p, x.(pqItem)) }
func (p *pq) Pop() any {
	old := *p
	v := old[len(old)-1]
	*p = old[:len(old)-1]
	return v
}

// PrimMST returns the total weight of a minimum spanning tree, growing one
// tree from vertex 0 and always taking the lightest edge leaving it.
// ok is false when the graph is disconnected (no spanning tree exists).
func PrimMST(n int, edges []Edge) (total int, ok bool) {
	if n == 0 {
		return 0, true
	}
	type arc struct{ to, w int }
	adj := make([][]arc, n)
	for _, e := range edges {
		adj[e.U] = append(adj[e.U], arc{e.V, e.W})
		adj[e.V] = append(adj[e.V], arc{e.U, e.W})
	}
	inTree := make([]bool, n)
	h := &pq{{0, 0}}
	count := 0
	for h.Len() > 0 && count < n {
		it := heap.Pop(h).(pqItem)
		if inTree[it.v] {
			continue
		}
		inTree[it.v] = true
		total += it.w
		count++
		for _, a := range adj[it.v] {
			if !inTree[a.to] {
				heap.Push(h, pqItem{a.w, a.to})
			}
		}
	}
	return total, count == n
}

func main() {
	edges := []Edge{{0, 1, 2}, {0, 2, 1}, {1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {4, 2, 2}}
	fmt.Println(PrimMST(5, edges))
}
