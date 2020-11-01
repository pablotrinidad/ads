package ads

import (
	"testing"
)

type fwtOpType uint

const (
	fwtOpAdd fwtOpType = iota
	fwtOpGet
)

type fwtOp struct {
	op    fwtOpType
	value int
	index int
	want  int
}

func TestFenwickTree(t *testing.T) {
	tests := []struct {
		name string
		size int
		ops  []fwtOp
	}{
		{
			name: "basic",
			size: 5,
			ops: []fwtOp{
				{op: fwtOpAdd, value: 1, index: 1},
				{op: fwtOpAdd, value: 2, index: 2},
				{op: fwtOpAdd, value: 3, index: 3},
				{op: fwtOpAdd, value: 4, index: 4},
				{op: fwtOpAdd, value: 5, index: 5},
				{op: fwtOpGet, index: 1, want: 1},
				{op: fwtOpGet, index: 2, want: 3},
				{op: fwtOpGet, index: 3, want: 6},
				{op: fwtOpGet, index: 4, want: 10},
				{op: fwtOpGet, index: 5, want: 15},
			},
		},
		{
			name: "updates and queries",
			size: 10,
			ops: []fwtOp{
				{op: fwtOpAdd, value: 234, index: 1},
				{op: fwtOpAdd, value: 6, index: 2},
				{op: fwtOpAdd, value: 1, index: 3},
				{op: fwtOpAdd, value: -61, index: 4},
				{op: fwtOpAdd, value: 78, index: 5},
				{op: fwtOpAdd, value: 69, index: 6},
				{op: fwtOpAdd, value: 72, index: 7},
				{op: fwtOpAdd, value: 11, index: 8},
				{op: fwtOpAdd, value: -21, index: 9},
				{op: fwtOpAdd, value: 25, index: 10},
				{op: fwtOpGet, index: 9, want: 389},
				{op: fwtOpAdd, value: 90, index: 10},
				{op: fwtOpAdd, value: -32, index: 8},
				{op: fwtOpGet, index: 2, want: 240},
				{op: fwtOpGet, index: 5, want: 258},
				{op: fwtOpGet, index: 8, want: 378},
				{op: fwtOpGet, index: 10, want: 472},
				{op: fwtOpAdd, value: -472, index: 1},
				{op: fwtOpGet, index: 10, want: 0},
				{op: fwtOpGet, index: 6, want: -145},
				{op: fwtOpAdd, value: 145, index: 5},
				{op: fwtOpGet, index: 6, want: 0},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fwt := NewFenwickTree(test.size)
			for _, op := range test.ops {
				switch op.op {
				case fwtOpAdd:
					if err := fwt.Add(op.index, op.value); err != nil {
						t.Fatalf("Add(%d, %d) returned unexpected error; %v", op.index, op.value, err)
					}
				case fwtOpGet:
					got, err := fwt.Get(op.index)
					if err != nil {
						t.Fatalf("Get(%d) returned unexpected error; %v", op.index, err)
					}
					if op.want != got {
						t.Errorf("Get(%d): %d, want %d", op.index, got, op.want)
					}
				}
			}
		})
	}
}

func TestFenwickTree_Errors(t *testing.T) {
	tests := []struct {
		name string
		size int
		ops  []fwtOp
	}{
		{
			name: "operating on empty tree",
			size: 0,
			ops: []fwtOp{
				{op: fwtOpGet, index: 1},
				{op: fwtOpAdd, index: 1, value: 1},
			},
		},
		{
			name: "invalid indexes",
			size: 10,
			ops: []fwtOp{
				{op: fwtOpGet, index: 0}, {op: fwtOpGet, index: 11},
				{op: fwtOpAdd, index: 0}, {op: fwtOpAdd, index: 11},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fwt := NewFenwickTree(test.size)
			for _, op := range test.ops {
				switch op.op {
				case fwtOpAdd:
					if err := fwt.Add(op.index, op.value); err == nil {
						t.Errorf("Add(%d, %d) returned non-nil error, want error", op.index, op.value)
					}
				case fwtOpGet:
					if _, err := fwt.Get(op.index); err == nil {
						t.Errorf("Get(%d) returned non-nil error, want error", op.index)
					}
				}
			}
		})
	}
}
