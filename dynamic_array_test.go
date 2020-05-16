package ads

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	arrayUnxOpt         = cmp.AllowUnexported(Array{})
	arrayIterableUnxOpt = cmp.AllowUnexported(ArrayIterable{})
)

func compareArrays(t *testing.T, want, got *Array, producer string) {
	if diff := cmp.Diff(want, got, arrayUnxOpt); diff != "" {
		t.Fatalf("%s() produced unwanted array: %s\nwant%s\ndiff want -> got\n%s",
			producer, got, want, diff)
	}
}

func TestArray_NewArray(t *testing.T) {
	want := Array{length: 0, capacity: 0, data: nil}
	got := NewArray()
	compareArrays(t, &want, &got, "NewArray")
}

func TestArray_Add(t *testing.T) {
	tests := []struct {
		name       string
		init, want Array
		insertData []int
	}{
		{
			name: "add to empty array",
			init: Array{length: 0, capacity: 0, data: []interface{}{}},
			want: Array{
				length:   1,
				capacity: arrayDefaultCapacity,
				data:     []interface{}{1, nil, nil, nil}},
			insertData: []int{1},
		},
		{
			name:       "add without resizing",
			init:       Array{length: 4, capacity: 5, data: []interface{}{1, 2, 3, 4, nil}},
			want:       Array{length: 5, capacity: 5, data: []interface{}{1, 2, 3, 4, 5}},
			insertData: []int{5},
		},
		{
			name: "resize once",
			init: Array{length: 32, capacity: 32, data: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
				13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}},
			want: Array{length: 33, capacity: 40, data: []interface{}{1, 2, 3, 4, 5, 6,
				7, 8, 9, 10, 11, 12,
				13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				33, nil, nil, nil, nil, nil, nil, nil}},
			insertData: []int{33},
		},
		{
			name: "resize multiple times",
			init: Array{},
			// Growth pattern is: 4, 8, 12, 16, 24, 32, 40, 48, 60, 72.
			want: Array{length: 65, capacity: 72, data: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9,
				10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
				30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
				50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, nil, nil, nil,
				nil, nil, nil, nil}},
			insertData: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, i := range test.insertData {
				test.init.Add(i)
			}
			compareArrays(t, &test.want, &test.init, "Add")
		})
	}
}

func TestArray_Remove(t *testing.T) {
	tests := []struct {
		name          string
		init, want    Array
		removeElement int
	}{
		{
			name:          "remove one in front",
			init:          Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want:          Array{length: 3, capacity: 4, data: []interface{}{2, 3, 4, nil}},
			removeElement: 1,
		},
		{
			name:          "remove one at the end",
			init:          Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want:          Array{length: 3, capacity: 4, data: []interface{}{1, 2, 3, nil}},
			removeElement: 4,
		},
		{
			name:          "remove one in the middle",
			init:          Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want:          Array{length: 3, capacity: 4, data: []interface{}{1, 3, 4, nil}},
			removeElement: 2,
		},
		{
			name:          "remove all",
			init:          Array{length: 4, capacity: 4, data: []interface{}{1, 1, 1, 1}},
			want:          Array{length: 0, capacity: 4, data: []interface{}{nil, nil, nil, nil}},
			removeElement: 1,
		},
		{
			name: "remove some",
			init: Array{length: 10, capacity: 10, data: []interface{}{1, 2, 3, 3, 4, 5,
				6, 3, 10, 3}},
			want: Array{length: 6, capacity: 10, data: []interface{}{1, 2, 4, 5, 6, 10,
				nil, nil, nil, nil}},
			removeElement: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.Remove(test.removeElement)
			compareArrays(t, &test.want, &test.init, "Remove")
		})
	}
}

func TestArray_RemoveIth(t *testing.T) {
	tests := []struct {
		name       string
		init, want Array
		i          int
	}{
		{
			name: "remove first",
			init: Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want: Array{length: 3, capacity: 4, data: []interface{}{2, 3, 4, nil}},
			i:    0,
		},
		{
			name: "remove last",
			init: Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want: Array{length: 3, capacity: 4, data: []interface{}{1, 2, 3, nil}},
			i:    3,
		},
		{
			name: "remove middle",
			init: Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want: Array{length: 3, capacity: 4, data: []interface{}{1, 3, 4, nil}},
			i:    1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.RemoveIth(test.i)
			compareArrays(t, &test.want, &test.init, "RemoveIth")
		})
	}
}

func TestArray_RemoveIthError(t *testing.T) {
	tests := []struct {
		name string
		a    Array
		i    int
	}{
		{
			name: "negative index",
			i:    -1,
		},
		{
			name: "zero index on empty init",
			i:    0,
		},
		{
			name: "zero index on empty init but with capacity",
			a:    Array{capacity: 10},
			i:    0,
		},
		{
			name: "index outside length",
			a:    Array{length: 2, capacity: 4, data: []interface{}{1, 2, nil, nil}},
			i:    2,
		},
		{
			name: "index outside capacity",
			a:    Array{length: 2, capacity: 2, data: []interface{}{1, 2}},
			i:    2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.a.RemoveIth(test.i); err == nil {
				t.Fatalf("RemoveIth(%d) returned non-nil error, want index error, init: %s",
					test.i, test.a)
			}
		})
	}
}

func TestArray_Get(t *testing.T) {
	tests := []struct {
		name     string
		a        Array
		i, v     int
		mustFail bool
	}{
		{
			name: "get first",
			a:    Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:    0, v: 1,
		},
		{
			name: "get midle",
			a:    Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:    2, v: 3,
		},
		{
			name: "get last",
			a:    Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:    3, v: 4,
		},
		{
			name:     "negative index",
			a:        Array{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:        -1,
			mustFail: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v, err := test.a.Get(test.i)
			if err == nil && test.mustFail {
				t.Fatalf("Get(%d) returned non-nil error, want index error, init: %s",
					test.i, test.a)
			}
			if v != test.v && !test.mustFail {
				t.Errorf("Get(%d) = %d, want %d", test.i, v, test.v)
			}
		})
	}
}

func TestArray_Contains(t *testing.T) {
	a := Array{
		length:   11,
		capacity: 32,
		data: []interface{}{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, nil, nil, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			nil},
	}
	var count int
	for i := 0; i <= 1024; i++ {
		contained := a.Contains(i)
		if contained {
			count++
			b := math.Log2(float64(i))
			if math.Round(b) != b {
				t.Errorf("Contains(%d) = true, but element is not in the array %s", i, a)
			}
		}
	}
	if count != a.Size() {
		t.Errorf("Got %d elements, must be all (%d)", count, a.Size())
	}
}

func TestArray_Size(t *testing.T) {
	var wantSize int
	a := NewArray()
	for i := 0; i <= 1024; i++ {
		a.Add(i)
		wantSize++
	}
	if wantSize != a.Size() {
		t.Errorf("Size() = %d, want %d", a.Size(), wantSize)
	}
	for i := 0; i <= 10; i++ {
		a.Remove(int(math.Pow(2, float64(i))))
		wantSize--
	}
	if wantSize != a.Size() {
		t.Errorf("Size() = %d, want %d", a.Size(), wantSize)
	}
}

func TestArray_Empty(t *testing.T) {
	ref := NewArray()
	got := NewArray()
	for i := 0; i < 100; i++ {
		got.Add(i)
	}
	got.Empty()
	compareArrays(t, &ref, &got, "Empty()")
}

func TestArray_String(t *testing.T) {
	tests := []struct {
		name, want string
		a          Array
	}{
		{
			name: "empty array",
			want: "[]",
			a:    NewArray(),
		},
		{
			name: "one element array",
			want: "[1]",
			a:    Array{length: 1, capacity: 1, data: []interface{}{1}},
		},
		{
			name: "multiple elements",
			want: "[1, 2, 3]",
			a:    Array{length: 3, capacity: 3, data: []interface{}{1, 2, 3}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.a.String()
			if test.want != got {
				t.Errorf("String() = %s, want %s", got, test.want)
			}
		})
	}
}

func TestArray_Iterator(t *testing.T) {
	arr := Array{
		length:   4,
		capacity: 6,
		data:     []interface{}{1, 2, 3, 4, nil, nil, nil, nil},
	}
	want := &ArrayIterable{i: 0, a: &arr}
	got := arr.Iterator()
	if diff := cmp.Diff(want, got, arrayUnxOpt, arrayIterableUnxOpt); diff != "" {
		t.Fatalf("Iterator() produced unwanted result: %v\nwant%v\ndiff want -> got\n%s",
			got, want, diff)
	}
}

func TestArrayIterable(t *testing.T) {
	arr := NewArray()
	n := 1000
	want := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr.Add(i)
		want = append(want, i)
	}
	i := arr.Iterator()
	got := make([]int, 0, arr.Size())
	for i.Scan() {
		v, err := i.Next()
		if err != nil {
			t.Fatalf("i.Next() got unexpected error %v", err)
		}
		got = append(got, v.(int))
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Iterator returned unexpected result: diff want -> got\n%s", diff)
	}
}

func TestArrayIterable_Error(t *testing.T) {
	arr := NewArray()
	arr.Add(1)
	arr.Add(2)
	i := arr.Iterator()
	for i.Scan() {
		i.i = -1
		if _, err := i.Next(); err == nil {
			t.Errorf("i.Next() returned non-nill error, want index error")
		}
		break
	}
}
