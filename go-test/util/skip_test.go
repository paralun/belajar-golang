package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkip(t *testing.T)  {
	t.Skip("Testing di skip")

	result := tambah(20, 20)
	assert.Equal(t, 40, result, "Result harusnya 40")
}
