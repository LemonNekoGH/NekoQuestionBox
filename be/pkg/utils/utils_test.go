package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexOf(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)
		r.Equal(2, IndexOf([]int{0, 1, 2, 3}, 2))
	})

	t.Run("not found", func(t *testing.T) {
		r := require.New(t)
		r.Equal(-1, IndexOf([]int{0, 1, 2, 3}, 4))
	})
}

func TestArrayContains(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)
		r.True(IsArrayContains([]int{0, 1, 2, 3}, 2))
	})

	t.Run("not found", func(t *testing.T) {
		r := require.New(t)
		r.False(IsArrayContains([]int{0, 1, 2, 3}, 4))
	})
}
