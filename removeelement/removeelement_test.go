package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testcases := []struct {
		in     []int
		remove int
		want   []int
	}{
		{[]int{1, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2, 3}, 1, []int{2, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2}},
		{[]int{}, 2, []int{}},
		{[]int{1, 2, 3, 2}, 2, []int{1, 3}},
		{[]int{1, 2, 3, 2}, 4, []int{1, 2, 3, 2}},
	}
	for _, tc := range testcases {
		result := removeElement(tc.in, tc.remove)
		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("RemoveElement, got: %v, want %v", result, tc.want)
		}
	}
}
