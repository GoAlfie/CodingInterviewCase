package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateItem(t *testing.T) {
	res := removeDuplicateItem([]any{1, 2, 2, 3, 3, 4})
	assert.Equal(t, res, []any{1, 2, 3, 4})
}

func TestPivotIndex(t *testing.T) {
	res := pivotIndex([]int{1, 7, 3, 7, 5, 6})
	assert.Equal(t, res, 3)
	res = pivotIndex([]int{9, -5, 2, 3})
	assert.Equal(t, res, 0)
	res = pivotIndex([]int{1, 2, 4})
	assert.Equal(t, res, -1)
}
