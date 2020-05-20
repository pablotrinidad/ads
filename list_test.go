package ads

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	listUnxOpt     = cmp.AllowUnexported(List{})
	listItemUnxOpt = cmp.AllowUnexported(ListItem{})
)

func compareLists(t *testing.T, want, got *List, producer string) {
	if diff := cmp.Diff(want, got, listUnxOpt, listItemUnxOpt); diff != "" {
		t.Fatalf("%s() produced unwanted list: %s\nwant%s\ndiff want -> got\n%s",
			producer, got, want, diff)
	}
}

func compareListItems(t *testing.T, want, got *ListItem, producer string) {
	if diff := cmp.Diff(want, got, listItemUnxOpt); diff != "" {
		t.Fatalf("%s() produced unwanted list item: %s\nwant%s\ndiff want -> got\n%s",
			producer, got, want, diff)
	}
}

func TestList_NewList(t *testing.T) {
	root := ListItem{}
	root.next = &root
	root.prev = &root
	want := &List{length: 0, root: root}
	got := NewList()
	compareLists(t, want, got, "NewList")
}

func TestListItem_Next(t *testing.T) {
	tmpList := NewList()
	tests := []struct {
		name string
		i    *ListItem
		want *ListItem
	}{
		{
			name: "nil values",
			i:    &ListItem{},
			want: nil,
		},
		{
			name: "with next item but without list",
			i:    &ListItem{next: &ListItem{Value: 1}},
			want: nil,
		},
		{
			name: "with list but without next item",
			i:    &ListItem{list: tmpList},
			want: nil,
		},
		{
			name: "with next item and underlying list",
			i:    &ListItem{next: &ListItem{Value: 1}, list: tmpList},
			want: &ListItem{Value: 1},
		},
		{
			name: "with next item and underlying list but pointing to root",
			i:    &ListItem{next: &tmpList.root, list: tmpList},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.i.Next()
			compareListItems(t, test.want, got, "Next")
		})
	}
}

func TestListItem_Prev(t *testing.T) {
	tmpList := NewList()
	tests := []struct {
		name string
		i    *ListItem
		want *ListItem
	}{
		{
			name: "nil values",
			i:    &ListItem{},
			want: nil,
		},
		{
			name: "with prev item but without list",
			i:    &ListItem{prev: &ListItem{Value: 1}},
			want: nil,
		},
		{
			name: "with list but without prev item",
			i:    &ListItem{list: tmpList},
			want: nil,
		},
		{
			name: "with prev item and underlying list",
			i:    &ListItem{prev: &ListItem{Value: 1}, list: tmpList},
			want: &ListItem{Value: 1},
		},
		{
			name: "with prev item and underlying list but pointing to root",
			i:    &ListItem{prev: &tmpList.root, list: tmpList},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.i.Prev()
			compareListItems(t, test.want, got, "Prev")
		})
	}
}

func TestListItem_String(t *testing.T) {
	tests := []struct {
		name, want string
		i          *ListItem
	}{
		{name: "sample 1", i: &ListItem{Value: 1}, want: "1"},
		{name: "sample 2", i: &ListItem{Value: []int{1, 2, 3}}, want: "[1 2 3]"},
		{name: "sample 3", i: &ListItem{Value: true}, want: "true"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.want != test.i.String() {
				t.Errorf("String(%v) = %v, want %v", test.i, test.i.String(), test.want)
			}
		})
	}
}

func ATestList_String(t *testing.T) {
	tests := []struct {
		name, want string
		content    []int
	}{
		{
			name:    "empty list",
			want:    "",
			content: make([]int, 0),
		},
		{
			name:    "one element array",
			want:    "1",
			content: []int{1},
		},
		{
			name:    "multiple elements",
			want:    "1 ↔ 2 ↔ 3",
			content: []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := "NewList()"
			if test.want != got {
				t.Errorf("String() = %s, want %s", got, test.want)
			}
		})
	}
}
