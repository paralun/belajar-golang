package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableHitung(t *testing.T)  {
	tests := []struct{
		nama string
		num1 int
		num2 int
		expected int
	}{
		{
			nama: "pertama",
			num1: 10,
			num2: 30,
			expected: 40,
		},
		{
			nama: "kedua",
			num1: 50,
			num2: 50,
			expected: 100,
		},
	}

	for _, test := range tests {
		t.Run(test.nama, func(t *testing.T) {
			result := tambah(test.num1, test.num2)
			assert.Equal(t, test.expected, result)
		})
	}
}
