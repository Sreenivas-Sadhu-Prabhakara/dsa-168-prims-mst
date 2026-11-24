# Minimum Spanning Tree (Prim)

> Striver SDE Sheet 168 · Day 24 Graph II — grow one tree, always absorbing the lightest edge that leaves it; a min-heap of frontier edges serves that choice. Go, stdlib only, tested.

## Problem

The minimum total weight connecting all vertices of a weighted undirected graph (the MST), with an honest failure on disconnected input.

## Approach

Prim grows a single tree from an arbitrary seed. The cut property does all the proving: for any cut of the graph, the lightest edge crossing it belongs to some MST — and "edges leaving the current tree" is exactly such a cut, so greedily absorbing the lightest frontier edge is always safe. The heap holds candidate (weight, vertex) entries; stale ones (vertex already absorbed) are skipped on pop, the same lazy deletion as Dijkstra — the two algorithms are structurally twins, differing only in what the heap key means (edge weight here, path distance there). Counting absorbed vertices doubles as the connectivity check.

**Complexity:** O(E log E) time, O(V+E) space.

## Visual walkthrough

Read the picture top to bottom — each numbered frame is one moment of the algorithm
running on this repo's test example.

![Step-by-step visual walkthrough](viz.svg)

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
