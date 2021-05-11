package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestTambahAssert(t *testing.T)  {
	result := tambah(20, 20)
	assert.Equal(t, 40, result, "Result harusnya 40")
}

func TestTambahAssert2(t *testing.T)  {
	result := tambah(20, 20)
	assert.NotEqual(t, 40, result, "Result harusnya 40")
}
