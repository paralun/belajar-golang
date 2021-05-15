package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubTest(t *testing.T)  {
	t.Run("first", func(t *testing.T) {
		result := tambah(20, 20)
		assert.Equal(t, 40, result, "Result harusnya 40")
	})

	t.Run("second", func(t *testing.T) {
		result := tambah(30, 20)
		assert.Equal(t, 50, result, "Result harusnya 50")
	})
}
