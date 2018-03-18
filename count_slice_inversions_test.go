package count_slice_inversions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCountSliceInversions(t *testing.T) {

	testSets := map[int][]int{
		45: {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		0:  {1, 2, 3, 4, 5, 6, 7, 8},
		1:  {2, 1, 3, 4, 5, 6, 7, 8},
		5:  {1, 5, 3, 4, 2, 6, 7, 8},
		14: {6, 5, 3, 4, 2, 1, 7, 8},
		3:  {4, 1, 2, 3, 5, 6, 7, 8, 9, 10},
		4:  {5, 1, 2, 3, 4, 6, 7, 8, 9, 10},
	}

	for k, v := range testSets {
		_, totalInv := CountSliceInversions(v)
		assert.Equal(t, k, totalInv)
	}
}
