package ads

import (
	"crypto/rand"
	"fmt"
	"testing"
)

type hashTableOpType int

const (
	hashTableSet hashTableOpType = iota
	hashTableGet
	hashTableDelete
	hashTableClear
)

type hashTableOp struct {
	op    hashTableOpType
	key   string
	value int
}

func TestHashTable_ThroughOps(t *testing.T) {
	tests := []struct {
		name string
		ops  []hashTableOp
	}{
		{
			name: "fill before resize then query",
			ops: []hashTableOp{
				{op: hashTableSet, key: "a", value: 1},
				{op: hashTableSet, key: "b", value: 2},
				{op: hashTableSet, key: "c", value: 3},
				{op: hashTableSet, key: "d", value: 4},
				{op: hashTableGet, key: "a"},
				{op: hashTableGet, key: "b"},
				{op: hashTableGet, key: "c"},
				{op: hashTableGet, key: "d"},
			},
		},
		{
			name: "fill and flush before resize then query",
			ops: []hashTableOp{
				{op: hashTableSet, key: "a", value: 1},
				{op: hashTableSet, key: "b", value: 2},
				{op: hashTableSet, key: "c", value: 3},
				{op: hashTableSet, key: "d", value: 4},
				{op: hashTableGet, key: "a"},
				{op: hashTableGet, key: "b"},
				{op: hashTableGet, key: "c"},
				{op: hashTableGet, key: "d"},
				{op: hashTableDelete, key: "a"},
				{op: hashTableDelete, key: "b"},
				{op: hashTableDelete, key: "c"},
				{op: hashTableDelete, key: "c"},
			},
		},
		{
			name: "fill and resize then query and empty",
			ops: []hashTableOp{
				{op: hashTableSet, key: "a", value: 1},
				{op: hashTableSet, key: "b", value: 2},
				{op: hashTableSet, key: "c", value: 3},
				{op: hashTableSet, key: "d", value: 4},
				{op: hashTableSet, key: "e", value: 5},
				{op: hashTableSet, key: "f", value: 6},
				{op: hashTableSet, key: "g", value: 7},
				{op: hashTableSet, key: "h", value: 8},
				{op: hashTableSet, key: "i", value: 9},
				{op: hashTableSet, key: "j", value: 10},
				{op: hashTableSet, key: "k", value: 11},
				{op: hashTableSet, key: "l", value: 12},
				{op: hashTableSet, key: "m", value: 13},
				{op: hashTableSet, key: "n", value: 14},
				{op: hashTableSet, key: "o", value: 15},
				{op: hashTableSet, key: "p", value: 16},
				{op: hashTableSet, key: "q", value: 17},
				{op: hashTableSet, key: "r", value: 18},
				{op: hashTableGet, key: "a"},
				{op: hashTableGet, key: "b"},
				{op: hashTableGet, key: "c"},
				{op: hashTableGet, key: "d"},
				{op: hashTableGet, key: "e"},
				{op: hashTableGet, key: "f"},
				{op: hashTableGet, key: "g"},
				{op: hashTableGet, key: "h"},
				{op: hashTableGet, key: "i"},
				{op: hashTableGet, key: "j"},
				{op: hashTableGet, key: "k"},
				{op: hashTableGet, key: "l"},
				{op: hashTableGet, key: "m"},
				{op: hashTableGet, key: "n"},
				{op: hashTableGet, key: "o"},
				{op: hashTableGet, key: "p"},
				{op: hashTableGet, key: "q"},
				{op: hashTableGet, key: "r"},
				{op: hashTableClear},
			},
		},
		{
			name: "query invalid keys",
			ops: []hashTableOp{
				{op: hashTableSet, key: "a", value: 1},
				{op: hashTableGet, key: "a"},
				{op: hashTableDelete, key: "a"},
				{op: hashTableGet, key: "a"},
				{op: hashTableSet, key: "b", value: 2},
				{op: hashTableGet, key: "b"},
				{op: hashTableDelete, key: "b"},
				{op: hashTableGet, key: "a"},
				{op: hashTableSet, key: "c", value: 3},
				{op: hashTableGet, key: "c"},
				{op: hashTableDelete, key: "c"},
				{op: hashTableGet, key: "c"},
				{op: hashTableSet, key: "d", value: 4},
				{op: hashTableGet, key: "d"},
				{op: hashTableDelete, key: "d"},
				{op: hashTableGet, key: "d"},
			},
		},
		{
			name: "reuse deleted buckets",
			ops: []hashTableOp{
				{op: hashTableSet, key: "a", value: 1},
				{op: hashTableGet, key: "a"},
				{op: hashTableDelete, key: "a"},
				{op: hashTableGet, key: "a"},
				{op: hashTableSet, key: "a", value: 1},
				{op: hashTableGet, key: "a"},
				{op: hashTableSet, key: "b", value: 2},
				{op: hashTableGet, key: "b"},
				{op: hashTableDelete, key: "b"},
				{op: hashTableSet, key: "b", value: 2},
				{op: hashTableGet, key: "b"},
				{op: hashTableGet, key: "a"},
				{op: hashTableSet, key: "c", value: 3},
				{op: hashTableGet, key: "c"},
				{op: hashTableDelete, key: "c"},
				{op: hashTableGet, key: "c"},
				{op: hashTableSet, key: "c", value: 3},
				{op: hashTableGet, key: "c"},
				{op: hashTableSet, key: "d", value: 4},
				{op: hashTableGet, key: "d"},
				{op: hashTableDelete, key: "d"},
				{op: hashTableGet, key: "d"},
				{op: hashTableSet, key: "d", value: 4},
				{op: hashTableGet, key: "d"},
			},
		},
		{
			name: "update values",
			ops: []hashTableOp{
				{op: hashTableSet, key: "a", value: 1},
				{op: hashTableGet, key: "a"},
				{op: hashTableSet, key: "a", value: 2},
				{op: hashTableSet, key: "b", value: 2},
				{op: hashTableGet, key: "b"},
				{op: hashTableSet, key: "b", value: 3},
				{op: hashTableSet, key: "c", value: 3},
				{op: hashTableGet, key: "c"},
				{op: hashTableSet, key: "c", value: 4},
				{op: hashTableSet, key: "d", value: 4},
				{op: hashTableGet, key: "d"},
				{op: hashTableSet, key: "d", value: 5},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			table := NewHashTable()
			control := make(map[string]int)

			for _, op := range test.ops {
				switch op.op {
				case hashTableSet:
					table.Set(op.key, op.value)
					control[op.key] = op.value
				case hashTableGet:
					tv, tok := table.Get(op.key)
					cv, cok := control[op.key]
					if (cok && tv.(int) != cv) || tok != cok {
						t.Fatalf("table.Get(%s): %d, %v want %d, %v", op.key, tv.(int), tok, cv, cok)
					}
				case hashTableDelete:
					table.Remove(op.key)
					delete(control, op.key)
				case hashTableClear:
					control = make(map[string]int)
					table.Empty()
				}
				if len(control) != table.Size() {
					t.Fatalf("table.Size(): %d, want %d", table.Size(), len(control))
				}
			}
		})
	}
}

func TestHashTable_LargeDataSets(t *testing.T) {
	tests := []int{
		4,
		5,
		50,
		100,
		1000,
		10000,
		100000,
	}

	genKey := func(n int) string {
		b := make([]byte, n)
		rand.Read(b)
		return fmt.Sprintf("%x", b)
	}

	for _, n := range tests {
		t.Run(fmt.Sprintf("n = %d", n), func(t *testing.T) {
			table := &HashTable{}
			control := make(map[string]int)

			// 1. Load n keys and query them
			for i := 1; i <= n; i++ {
				key := genKey(20)
				table.Set(key, i)
				control[key] = i
				tv, tok := table.Get(key)
				cv, cok := control[key]
				if tv != cv || !tok || !cok {
					t.Fatalf("table.Get(%s): %d, %v want %d, %v", key, tv.(int), tok, cv, cok)
				}
				if len(control) != table.Size() {
					t.Fatalf("table.Size(): %d, want %d", table.Size(), len(control))
				}
			}
			// 2. Update n keys and query
			for k, v := range control {
				table.Set(k, -1*v)
				control[k] = -1 * v
				tv, tok := table.Get(k)
				cv, cok := control[k]
				if tv != cv || !tok || !cok {
					t.Fatalf("table.Get(%s): %d, %v want %d, %v", k, tv.(int), tok, cv, cok)
				}
				if len(control) != table.Size() {
					t.Fatalf("table.Size(): %d, want %d", table.Size(), len(control))
				}
			}
			// 3. Query n / 3 random keys
			for i := 0; i < n/3; i++ {
				key := genKey(20)
				tv, tok := table.Get(key)
				cv, cok := control[key]
				if tok != cok || (cok && tv.(int) != cv) {
					t.Fatalf("table.Get(%s): %d, %v want %d, %v", key, tv.(int), tok, cv, cok)
				}
			}
			// 4. Delete n / 3 random keys
			count := 0
			for k := range control {
				if count > n/3 {
					break
				}
				table.Remove(k)
				delete(control, k)
				if len(control) != table.Size() {
					t.Fatalf("table.Size(): %d, want %d", table.Size(), len(control))
				}
			}
			// 5. Query all keys
			for k := range control {
				tv, tok := table.Get(k)
				cv, cok := control[k]
				if tv != cv || !tok || !cok {
					t.Fatalf("table.Get(%s): %d, %v want %d, %v", k, tv.(int), tok, cv, cok)
				}
				if len(control) != table.Size() {
					t.Fatalf("table.Size(): %d, want %d", table.Size(), len(control))
				}
			}
			// 6. Flush table
			control = make(map[string]int)
			table.Empty()
			if len(control) != table.Size() {
				t.Fatalf("table.Size(): %d, want %d", table.Size(), len(control))
			}
		})
	}
}
