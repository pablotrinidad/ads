package ads

import (
	"fmt"
	"testing"
)

func logQueueSatisfaction(t *testing.T, ds string, q Queue) {
	t.Helper()
	t.Logf("%s satisfies Queue interface: %v", ds, q)
}

// TestQueueInterfaceSatisfaction verifies (during compilation) that the
// multiple queue implementations satisfy the Queue interface.
func TestQueueInterfaceSatisfaction(t *testing.T) {
	var q Queue

	// Array-based
	q = NewArrayBasedQueue(1)
	logQueueSatisfaction(t, "ArrayBasedStack", q)
}

type queueOpType int

const (
	queuePush queueOpType = iota
	queuePop
	queueEmpty
)

type queueOp struct {
	op       queueOpType
	input    int
	mustFail bool
	want     int
}

type intErrorResult struct {
	v        int
	mustFail bool
}

func validateQueueError(t *testing.T, producer string, err error, mustFail bool) {
	t.Helper()
	if err != nil && !mustFail {
		t.Fatalf("%s failed; %v", producer, err)
	} else if err == nil && mustFail {
		t.Fatalf("%s returned non-nil error, want error", producer)
	}
}

// TestArrayBasedQueue_ThroughOps will perform a set of operations against a newly
// created queue and check for read-only operations results after each one is executed.
// It will also perform the read-only operations before executing the test case operations.
func TestArrayBasedQueue_ThroughOps(t *testing.T) {
	tests := []struct {
		name      string
		size      int
		ops       []queueOp
		wantFront []intErrorResult
		wantBack  []intErrorResult
		wantSize  []int
	}{
		{
			name: "fill and flush",
			size: 4,
			ops: []queueOp{
				{op: queuePush, input: 1},
				{op: queuePush, input: 2},
				{op: queuePush, input: 3},
				{op: queuePush, input: 4},
				{op: queuePop, want: 1},
				{op: queuePop, want: 2},
				{op: queuePop, want: 3},
				{op: queuePop, want: 4},
			},
			wantFront: []intErrorResult{
				{v: 1}, {v: 1}, {v: 1}, {v: 1},
				{v: 2}, {v: 3}, {v: 4}, {mustFail: true},
			},
			wantBack: []intErrorResult{
				{v: 1}, {v: 2}, {v: 3}, {v: 4},
				{v: 4}, {v: 4}, {v: 4}, {mustFail: true},
			},
			wantSize: []int{1, 2, 3, 4, 3, 2, 1, 0},
		},
		{
			name: "fill and empty",
			size: 8,
			ops: []queueOp{
				{op: queuePush, input: 1},
				{op: queuePush, input: 2},
				{op: queuePush, input: 3},
				{op: queuePush, input: 4},
				{op: queuePush, input: 5},
				{op: queuePush, input: 6},
				{op: queuePush, input: 7},
				{op: queuePush, input: 8},
				{op: queueEmpty},
			},
			wantFront: []intErrorResult{
				{v: 1}, {v: 1}, {v: 1}, {v: 1},
				{v: 1}, {v: 1}, {v: 1}, {v: 1},
				{mustFail: true},
			},
			wantBack: []intErrorResult{
				{v: 1}, {v: 2}, {v: 3}, {v: 4},
				{v: 5}, {v: 6}, {v: 7}, {v: 8},
				{mustFail: true},
			},
			wantSize: []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
		},
		{
			name: "cycle through underlying array one by one",
			size: 3,
			ops: []queueOp{
				{op: queuePush, input: 1}, {op: queuePop, want: 1},
				{op: queuePush, input: 1}, {op: queuePop, want: 1},
				{op: queuePush, input: 1}, {op: queuePop, want: 1},
				{op: queuePush, input: 1}, {op: queuePop, want: 1},
			},
			wantFront: []intErrorResult{
				{v: 1}, {mustFail: true},
				{v: 1}, {mustFail: true},
				{v: 1}, {mustFail: true},
				{v: 1}, {mustFail: true},
			},
			wantBack: []intErrorResult{
				{v: 1}, {mustFail: true},
				{v: 1}, {mustFail: true},
				{v: 1}, {mustFail: true},
				{v: 1}, {mustFail: true},
			},
			wantSize: []int{
				1, 0,
				1, 0,
				1, 0,
				1, 0,
			},
		},
		{
			name: "queue overflow",
			size: 3,
			ops: []queueOp{
				{op: queuePush, input: 1},
				{op: queuePush, input: 2},
				{op: queuePush, input: 3},
				{op: queuePush, input: 4, mustFail: true},
				{op: queuePop, want: 1},
				{op: queuePop, want: 2},
				{op: queuePop, want: 3},
				{op: queuePop, mustFail: true},
				{op: queuePush, input: 4},
				{op: queuePush, input: 5},
				{op: queuePush, input: 6},
				{op: queuePush, input: 7, mustFail: true},
				{op: queueEmpty},
			},
			wantFront: []intErrorResult{
				{v: 1}, {v: 1}, {v: 1}, {v: 1},
				{v: 2}, {v: 3}, {mustFail: true}, {mustFail: true},
				{v: 4}, {v: 4}, {v: 4}, {v: 4},
				{mustFail: true},
			},
			wantBack: []intErrorResult{
				{v: 1}, {v: 2}, {v: 3}, {v: 3},
				{v: 3}, {v: 3}, {mustFail: true}, {mustFail: true},
				{v: 4}, {v: 5}, {v: 6}, {v: 6},
				{mustFail: true},
			},
			wantSize: []int{
				1, 2, 3, 3,
				2, 1, 0, 0,
				1, 2, 3, 3,
				0,
			},
		},
		{
			name: "queue of size 0",
			size: 0,
			ops: []queueOp{
				{op: queuePush, mustFail: true},
				{op: queuePop, mustFail: true},
				{op: queueEmpty},
			},
			wantFront: []intErrorResult{
				{mustFail: true}, {mustFail: true}, {mustFail: true},
			},
			wantBack: []intErrorResult{
				{mustFail: true}, {mustFail: true}, {mustFail: true},
			},
			wantSize: []int{0, 0, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := NewArrayBasedQueue(uint(test.size))

			// Perform read-only operations once
			if _, err := q.Front(); err == nil {
				t.Errorf("Front() returned non-nill error on empty queue, want error")
			}
			if _, err := q.Back(); err == nil {
				t.Errorf("Back() returned non-nill error on empty queue, want error")
			}
			if s := q.Size(); s != 0 {
				t.Errorf("Size() = %d, want 0", s)
			}

			for i, op := range test.ops {
				switch op.op {
				case queuePush:
					validateQueueError(
						t, fmt.Sprintf("Push(%d)", op.input),
						q.Push(op.input),
						op.mustFail)
				case queuePop:
					popResult, popError := q.Pop()
					validateQueueError(t, "Pop()", popError, op.mustFail)
					if !op.mustFail && popResult.(int) != op.want {
						t.Fatalf("Pop(): %d, want %d", popResult, op.want)
					}
				case queueEmpty:
					q.Empty()
				}

				wantFront := test.wantFront[i]
				frontGot, frontErr := q.Front()
				validateQueueError(t, "Front()", frontErr, wantFront.mustFail)
				if !wantFront.mustFail && frontGot.(int) != wantFront.v {
					t.Errorf("Front(): %d, want %d", frontGot, wantFront.v)
				}

				wantBack := test.wantBack[i]
				backGot, backErr := q.Back()
				validateQueueError(t, "Back()", backErr, wantBack.mustFail)
				if !wantBack.mustFail && backGot.(int) != wantBack.v {
					t.Errorf("Back(): %d, want %d", backGot, wantBack.v)
				}

				if q.Size() != test.wantSize[i] {
					t.Errorf("Size(): %d, want %d", q.Size(), test.wantSize[i])
				}
			}
		})
	}
}
