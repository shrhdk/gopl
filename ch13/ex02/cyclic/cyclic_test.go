package cyclic

import "testing"

func TestIsCyclic(t *testing.T) {
	var item struct{}
	nonCyclicMap := make(map[int]interface{})
	nonCyclicMap[0] = &item
	nonCyclicMap[1] = &item

	type LinkedList struct{ next *LinkedList }
	var item1, item2 LinkedList
	item1.next = &item2
	item2.next = &item1

	cyclicStruct := struct{ i interface{} }{}
	cyclicStruct.i = &cyclicStruct

	var cyclicArray [1]interface{}
	cyclicArray[0] = &cyclicArray

	tests := []struct {
		give interface{}
		want bool
	}{
		{nil, false},
		{0, false},
		{1.0, false},
		{true, false},
		{false, false},
		{"", false},
		{[0]int{}, false},
		{make([]int, 0), false},
		{make(map[bool]bool), false},
		{&struct{}{}, false},
		{nonCyclicMap, false},
		{item1, true},
		{cyclicStruct, true},
		{cyclicArray, true},
	}

	for _, test := range tests {
		got := IsCyclic(test.give)
		if got != test.want {
			t.Errorf("give %s, want %v, got %v", test.give, test.want, got)
		}
	}
}
