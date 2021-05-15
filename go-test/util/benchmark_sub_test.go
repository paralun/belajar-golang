package util

import "testing"

func BenchmarkSub(b *testing.B)  {
	b.Run("pertama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tambah(20, 70)
		}
	})

	b.Run("kedua", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tambah(100, 200)
		}
	})
}
