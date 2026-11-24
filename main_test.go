package main

import "testing"

func TestPrimMST(t *testing.T) {
	edges := []Edge{{0, 1, 2}, {0, 2, 1}, {1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {4, 2, 2}}
	total, ok := PrimMST(5, edges)
	if !ok || total != 5 { // 1 + 1 + 2 + 1
		t.Errorf("PrimMST = (%d,%v), want 5", total, ok)
	}

	// Disconnected: no spanning tree.
	if _, ok := PrimMST(3, []Edge{{0, 1, 1}}); ok {
		t.Error("disconnected graph must report not-ok")
	}

	// Single vertex: empty tree of weight 0.
	if total, ok := PrimMST(1, nil); !ok || total != 0 {
		t.Errorf("single vertex = (%d,%v)", total, ok)
	}

	// Parallel edges: the lighter wins.
	total, ok = PrimMST(2, []Edge{{0, 1, 9}, {0, 1, 3}})
	if !ok || total != 3 {
		t.Errorf("parallel edges = (%d,%v), want 3", total, ok)
	}
}
