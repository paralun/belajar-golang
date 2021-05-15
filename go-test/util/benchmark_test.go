package util

import "testing"

func BenchmarkHitung1(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		tambah(20, 30)
	}
}

func BenchmarkHitung2(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		tambah(100, 200)
	}
}
