package ads

import "fmt"

// FenwickTree is a binary indexed tree with optimal complexity for computing prefix sums.
// Important: this implementation is one-indexed.
type FenwickTree struct {
	bit []int
	n   int
}

// NewFenwickTree returns a zero-initialized Fenwick Tree.
func NewFenwickTree(size int) *FenwickTree {
	return &FenwickTree{n: size, bit: make([]int, size+1)}
}

// Get returns the prefix sum of the first i elements.
func (t *FenwickTree) Get(i int) (int, error) {
	if i > t.n {
		return 0, fmt.Errorf("tree size is %d, got query for %d elements", t.n, i)
	}
	s := 0
	for i > 0 {
		s += t.bit[i]
		i -= i & (-i)
	}
	return s, nil
}

// Add value at the given index.
func (t *FenwickTree) Add(i, value int) error {
	if i <= 0 || i > t.n {
		return fmt.Errorf("invalid index %d", i)
	}
	fmt.Printf("Adding %d at index %d\n", value, i)
	for i <= t.n {
		fmt.Printf("\tindex %d\n", i)
		t.bit[i] += value
		i += i & (-i)
	}
	return nil
}
