package ads

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var unxOpt = cmp.AllowUnexported(List{})

func TestList_NewList(t *testing.T) {
	want := List{length: 0, capacity: 0, data: nil}
	got := NewList()
	if diff := cmp.Diff(want, got, unxOpt); diff != "" {
		t.Fatalf("NewList() = %v, want %v: diff want -> got\n%s", got, want, diff)
	}
}

func TestList_Add(t *testing.T) {
	tests := []struct {
		name       string
		init, want List
		insertData []int
	}{
		{
			name: "add to empty list",
			init: List{length: 0, capacity: 0, data: []interface{}{}},
			want: List{
				length:   1,
				capacity: listDefaultCapacity,
				data:     []interface{}{1, nil, nil, nil}},
			insertData: []int{1},
		},
		{
			name:       "add without resizing",
			init:       List{length: 4, capacity: 5, data: []interface{}{1, 2, 3, 4, nil}},
			want:       List{length: 5, capacity: 5, data: []interface{}{1, 2, 3, 4, 5}},
			insertData: []int{5},
		},
		{
			name: "resize once",
			init: List{length: 32, capacity: 32, data: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
				13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}},
			want: List{length: 33, capacity: 40, data: []interface{}{1, 2, 3, 4, 5, 6,
				7, 8, 9, 10, 11, 12,
				13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				33, nil, nil, nil, nil, nil, nil, nil}},
			insertData: []int{33},
		},
		{
			name: "resize multiple times",
			init: List{},
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
				test.init.Add(i)
			}
			if diff := cmp.Diff(test.init, test.want, unxOpt); diff != "" {
				t.Fatalf("Add() produced unwanted list: %s\nwant%s\ndiff want -> got\n%s",
					test.init, test.want, diff)
			}
		})
	}
}
