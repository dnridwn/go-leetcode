// https://leetcode.com/problems/merge-sorted-array/
package merge_sorted_array

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tc := []struct {
		nums1    []int
		nums2    []int
		m        int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			nums2:    []int{2, 5, 6},
			m:        3,
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			nums2:    []int{},
			m:        1,
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			nums2:    []int{1},
			m:        0,
			n:        1,
			expected: []int{1},
		},
	}

	for k, tt := range tc {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, true, reflect.DeepEqual(tt.nums1, tt.expected))
		})
	}
}
