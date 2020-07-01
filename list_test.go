package ads

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	listUnxOpt         = cmp.AllowUnexported(List{})
	listItemUnxOpt     = cmp.AllowUnexported(ListItem{})
	listIterableUnxOpt = cmp.AllowUnexported(ListIterable{})
)

func compareLists(t *testing.T, want, got *List, producer string) {
	if diff := cmp.Diff(want, got, listUnxOpt, listItemUnxOpt); diff != "" {
		t.Fatalf("%s() produced unwanted list: %s\nwant%s\ndiff want -> got\n%s",
			producer, got, want, diff)
	}
}

func compareListItems(t *testing.T, want, got *ListItem, producer string) {
	if diff := cmp.Diff(want, got, listItemUnxOpt, listUnxOpt); diff != "" {
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

func TestList_String(t *testing.T) {
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
			l := NewList()
			for _, v := range test.content {
				l.Add(v)
			}
			got := l.String()
			if test.want != got {
				t.Errorf("String() = %s, want %s", got, test.want)
			}
		})
	}
}

func TestList_Contains(t *testing.T) {
	n := 100
	l := NewList()
	for i := 1; i <= n; i++ {
		l.Add(i)
	}
	for i := 1; i <= n*2; i++ {
		if l.Contains(i) && i > 100 {
			t.Fatalf("Contains(%d) = false, want true", i)
		}
	}
}

func TestList_Iterator(t *testing.T) {
	l := NewList()
	want := &ListIterable{l: l, n: l.Head()}
	got := l.Iterator()
	if diff := cmp.Diff(want, got, listIterableUnxOpt, listUnxOpt, listItemUnxOpt); diff != "" {
		t.Fatalf("Iterator() produced unwanted result: %v\nwant%v\ndiff want -> got\n%s",
			got, want, diff)
	}
}

func TestListIterable(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{name: "empty list", n: 0},
		{name: "one element", n: 1},
		{name: "multiple elements", n: 10},
		{name: "more elements", n: 1000},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := List{}
			want := make([]int, test.n)
			for i := 0; i < test.n; i++ {
				l.Add(i)
				want[i] = i
			}
			got := make([]int, 0, l.Size())
			i := l.Iterator()
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
		})
	}
}

func TestListIterable_Error(t *testing.T) {
	l := NewList()
	i := l.Iterator()
	if _, err := i.Next(); err == nil {
		t.Errorf("i.Next() returned non-nill error, want error")
	}
}

func TestList_GetItem(t *testing.T) {
	tests := []struct {
		name      string
		content   []int
		q         int
		contained bool
	}{
		{name: "empty list"},
		{
			name:      "valid query in one-element list",
			content:   []int{1},
			q:         1,
			contained: true,
		},
		{
			name:      "invalid query in one-element list",
			content:   []int{1},
			q:         0,
			contained: false,
		},
		{
			name:      "valid query with value located at the beginning",
			content:   []int{1, 2, 3, 4, 5},
			q:         1,
			contained: true,
		},
		{
			name:      "valid query with value located at the end",
			content:   []int{1, 2, 3, 4, 5},
			q:         5,
			contained: true,
		},
		{
			name:      "valid query with value located in the middle",
			content:   []int{1, 2, 3, 4, 5},
			q:         3,
			contained: true,
		},
		{
			name:      "invalid query in long list",
			content:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			q:         16,
			contained: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := List{}
			var qFound bool
			var wantItem *ListItem
			for i, x := range test.content {
				l.Add(x)
				if test.contained && i == 0 {
					wantItem = l.Head()
				} else if test.contained && !qFound {
					wantItem = wantItem.Next()
				}
				qFound = qFound || x == test.q
			}
			if !qFound && test.contained {
				t.Fatalf("lost track of query value %d", test.q)
			}
			got, gotError := l.GetItem(test.q)
			if gotError == nil && !test.contained {
				t.Errorf("GetItem(%d) returned non-nill error, want error for missing value",
					test.q)
			}
			compareListItems(t, wantItem, got, "GetItem")
		})
	}
}

func TestList_RemoveItem(t *testing.T) {
	tests := []struct {
		name          string
		content, want []int
		deletionItem  int
	}{
		{name: "empty list"},
		{
			name:         "one-element list",
			content:      []int{1},
			want:         []int{},
			deletionItem: 1,
		},
		{
			name:         "one-element removing nil",
			content:      []int{1},
			want:         []int{1},
			deletionItem: 0,
		},
		{
			name:         "multiple elements removing first",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{2, 3, 4, 5},
			deletionItem: 1,
		},
		{
			name:         "multiple elements removing second",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{1, 3, 4, 5},
			deletionItem: 2,
		},
		{
			name:         "multiple elements removing last",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{1, 2, 3, 4},
			deletionItem: 5,
		},
		{
			name:         "multiple elements removing one before last",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{1, 2, 3, 5},
			deletionItem: 4,
		},
		{
			name:         "multiple elements removing middle",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{1, 2, 4, 5},
			deletionItem: 3,
		},
		{
			name:         "two elements removing last",
			content:      []int{1, 2},
			want:         []int{1},
			deletionItem: 2,
		},
		{
			name:         "two elements removing first",
			content:      []int{1, 2},
			want:         []int{2},
			deletionItem: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewList()
			var item *ListItem
			var err error
			for _, x := range test.content {
				got.Add(x)
				if x == test.deletionItem && item == nil {
					item, err = got.GetItem(x)
					if err != nil {
						t.Fatalf("failed getting element %d after insertion", x)
					}
				}
			}
			got.RemoveItem(item)
			want := NewList()
			for _, x := range test.want {
				want.Add(x)
			}
			compareLists(t, want, got, "RemoveItem")
		})
	}
}

func TestList_Remove(t *testing.T) {
	tests := []struct {
		name          string
		content, want []int
		deletionItem  int
	}{
		{name: "empty list"},
		{
			name:         "1 element list no deletions",
			content:      []int{1},
			want:         []int{1},
			deletionItem: 0,
		},
		{
			name:         "1 element list with deletions",
			content:      []int{1},
			want:         []int{},
			deletionItem: 1,
		},
		{
			name:         "multiple elements list without deletions",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{1, 2, 3, 4, 5},
			deletionItem: 0,
		},
		{
			name:         "multiple elements list with deletion at the beginning",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{2, 3, 4, 5},
			deletionItem: 1,
		},
		{
			name: "multiple elements list with deletion at the beginning repeated" +
				" contiguous",
			content:      []int{1, 1, 1, 4, 5},
			want:         []int{4, 5},
			deletionItem: 1,
		},
		{
			name:         "multiple elements list with deletion in the middle",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{1, 2, 4, 5},
			deletionItem: 3,
		},
		{
			name:         "multiple elements list with deletion in the middle repeated contiguous",
			content:      []int{1, 2, 2, 2, 5},
			want:         []int{1, 5},
			deletionItem: 2,
		},
		{
			name:         "multiple elements list with deletion at the end",
			content:      []int{1, 2, 3, 4, 5},
			want:         []int{1, 2, 3, 4},
			deletionItem: 5,
		},
		{
			name:         "multiple elements list with deletion at the end repeated contiguous",
			content:      []int{1, 2, 3, 3, 3},
			want:         []int{1, 2},
			deletionItem: 3,
		},
		{
			name:         "multiple elements all repeated",
			content:      []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:         []int{},
			deletionItem: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewList()
			for _, x := range test.content {
				got.Add(x)
			}
			got.Remove(test.deletionItem)
			want := NewList()
			for _, x := range test.want {
				want.Add(x)
			}
			compareLists(t, want, got, "Remove")
		})
	}
}

func TestList_HeadTail(t *testing.T) {
	tests := []struct {
		name               string
		content            []int
		wantHead, wantTail int
	}{
		{name: "empty list"},
		{name: "one-element list", content: []int{1}, wantHead: 1, wantTail: 1},
		{name: "two-element list", content: []int{1, 2}, wantHead: 1, wantTail: 2},
		{name: "three-element list", content: []int{1, 2, 3}, wantHead: 1, wantTail: 3},
		{
			name:     "multiple elements list",
			content:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			wantHead: 1,
			wantTail: 20,
		},
	}
	for _, test := range tests {
		l := NewList()
		for _, x := range test.content {
			l.Add(x)
		}
		t.Run(fmt.Sprintf("%s head", test.name), func(t *testing.T) {
			got := l.Head()
			if test.wantHead != 0 && got.Value != test.wantHead {
				t.Errorf("Head(): %d, want %d for list %s", got.Value, test.wantHead, l)
			}
		})
		t.Run(fmt.Sprintf("%s tail", test.name), func(t *testing.T) {
			got := l.Tail()
			if test.wantTail != 0 && got.Value != test.wantTail {
				t.Errorf("Tail(): %d, want %d for list %s", got.Value, test.wantTail, l)
			}
		})
	}
}

func TestList_Empty(t *testing.T) {
	tests := []struct {
		name    string
		content []int
	}{
		{name: "empty list"},
		{name: "one-element list", content: []int{1}},
		{name: "two-element list", content: []int{1, 2}},
		{name: "three-element list", content: []int{1, 2, 3}},
		{name: "multiple elements list", content: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	emptyList := NewList()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewList()
			for _, x := range test.content {
				got.Add(x)
			}
			got.Empty()
			compareLists(t, emptyList, got, "Empty")
		})
	}
}
