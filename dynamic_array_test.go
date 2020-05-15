package ads

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var unxOpt = cmp.AllowUnexported(List{})

func compareLists(t *testing.T, want, got *List, producer string) {
	if diff := cmp.Diff(want, got, unxOpt); diff != "" {
		t.Fatalf("%s() produced unwanted list: %s\nwant%s\ndiff want -> got\n%s",
			producer, want, got, diff)
	}
}

func TestList_NewList(t *testing.T) {
	want := List{length: 0, capacity: 0, data: nil}
	got := NewList()
	compareLists(t, &want, &got, "NewList")
}

func TestList_Add(t *testing.T) {
	tests := []struct {
		name       string
		list, want List
		insertData []int
	}{
		{
			name: "add to empty list",
			list: List{length: 0, capacity: 0, data: []interface{}{}},
			want: List{
				length:   1,
				capacity: listDefaultCapacity,
				data:     []interface{}{1, nil, nil, nil}},
			insertData: []int{1},
		},
		{
			name:       "add without resizing",
			list:       List{length: 4, capacity: 5, data: []interface{}{1, 2, 3, 4, nil}},
			want:       List{length: 5, capacity: 5, data: []interface{}{1, 2, 3, 4, 5}},
			insertData: []int{5},
		},
		{
			name: "resize once",
			list: List{length: 32, capacity: 32, data: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
				13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}},
			want: List{length: 33, capacity: 40, data: []interface{}{1, 2, 3, 4, 5, 6,
				7, 8, 9, 10, 11, 12,
				13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				33, nil, nil, nil, nil, nil, nil, nil}},
			insertData: []int{33},
		},
		{
			name: "resize multiple times",
			list: List{},
			// Growth pattern is: 4, 8, 12, 16, 24, 32, 40, 48, 60, 72.
			want: List{length: 65, capacity: 72, data: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9,
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
				test.list.Add(i)
			}
			compareLists(t, &test.want, &test.list, "Add")
		})
	}
}

func TestList_Remove(t *testing.T) {
	tests := []struct {
		name          string
		list, want    List
		removeElement int
	}{
		{
			name:          "remove one in front",
			list:          List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want:          List{length: 3, capacity: 4, data: []interface{}{2, 3, 4, nil}},
			removeElement: 1,
		},
		{
			name:          "remove one at the end",
			list:          List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want:          List{length: 3, capacity: 4, data: []interface{}{1, 2, 3, nil}},
			removeElement: 4,
		},
		{
			name:          "remove one in the middle",
			list:          List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want:          List{length: 3, capacity: 4, data: []interface{}{1, 3, 4, nil}},
			removeElement: 2,
		},
		{
			name:          "remove all",
			list:          List{length: 4, capacity: 4, data: []interface{}{1, 1, 1, 1}},
			want:          List{length: 0, capacity: 4, data: []interface{}{nil, nil, nil, nil}},
			removeElement: 1,
		},
		{
			name: "remove some",
			list: List{length: 10, capacity: 10, data: []interface{}{1, 2, 3, 3, 4, 5,
				6, 3, 10, 3}},
			want: List{length: 6, capacity: 10, data: []interface{}{1, 2, 4, 5, 6, 10,
				nil, nil, nil, nil}},
			removeElement: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.list.Remove(test.removeElement)
			compareLists(t, &test.want, &test.list, "Remove")
		})
	}
}

func TestList_RemoveIth(t *testing.T) {
	tests := []struct {
		name       string
		list, want List
		i          int
	}{
		{
			name: "remove first",
			list: List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want: List{length: 3, capacity: 4, data: []interface{}{2, 3, 4, nil}},
			i:    0,
		},
		{
			name: "remove last",
			list: List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want: List{length: 3, capacity: 4, data: []interface{}{1, 2, 3, nil}},
			i:    3,
		},
		{
			name: "remove middle",
			list: List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			want: List{length: 3, capacity: 4, data: []interface{}{1, 3, 4, nil}},
			i:    1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.list.RemoveIth(test.i)
			compareLists(t, &test.want, &test.list, "RemoveIth")
		})
	}
}

func TestList_RemoveIthError(t *testing.T) {
	tests := []struct {
		name string
		list List
		i    int
	}{
		{
			name: "negative index",
			i:    -1,
		},
		{
			name: "zero index on empty list",
			i:    0,
		},
		{
			name: "zero index on empty list but with capacity",
			list: List{capacity: 10},
			i:    0,
		},
		{
			name: "index outside length",
			list: List{length: 2, capacity: 4, data: []interface{}{1, 2, nil, nil}},
			i:    2,
		},
		{
			name: "index outside capacity",
			list: List{length: 2, capacity: 2, data: []interface{}{1, 2}},
			i:    2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.list.RemoveIth(test.i); err == nil {
				t.Fatalf("RemoveIth(%d) returned non-nil error, want index error, list: %s",
					test.i, test.list)
			}
		})
	}
}

func TestList_Get(t *testing.T) {
	tests := []struct {
		name     string
		list     List
		i, v     int
		mustFail bool
	}{
		{
			name: "get first",
			list: List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:    0, v: 1,
		},
		{
			name: "get midle",
			list: List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:    2, v: 3,
		},
		{
			name: "get last",
			list: List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:    3, v: 4,
		},
		{
			name:     "negative index",
			list:     List{length: 4, capacity: 4, data: []interface{}{1, 2, 3, 4}},
			i:        -1,
			mustFail: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v, err := test.list.Get(test.i)
			if err == nil && test.mustFail {
				t.Fatalf("Get(%d) returned non-nil error, want index error, list: %s",
					test.i, test.list)
			}
			if v != test.v && !test.mustFail {
				t.Errorf("Get(%d) = %d, want %d", test.i, v, test.v)
			}
		})
	}
}
