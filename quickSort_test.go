package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Case 1", []int{34, 7, 23, 32, 5, 62, 31}, []int{5, 7, 23, 31, 32, 34, 62}},
		{"Case 2", []int{5, 3, 8, 1, 2, 4}, []int{1, 2, 3, 4, 5, 8}},
		{"Empty Slice", []int{}, []int{}},
		{"Single Element", []int{1}, []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := QuickSort(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("QuickSort(%v) = %v; expected %v", test.input, result, test.expected)
			}
		})
	}
}
