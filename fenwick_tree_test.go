package ads

import "testing"

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
