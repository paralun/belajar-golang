package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// Require memanggil function FailNow()
func TestTambahRequire(t *testing.T)  {
	result := tambah(20, 20)
	require.Equal(t, 30, result, "Result harusnya 30")
	fmt.Println("Testing Selesai")
}