package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M)  {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//After
	fmt.Println("AFTER UNIT TEST")
}

func TestBeforeAfter(t *testing.T)  {
	result := tambah(20, 20)
	assert.Equal(t, 40, result, "Result harusnya 40")
}
