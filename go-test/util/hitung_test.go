package util

import (
	"fmt"
	"testing"
)

func TestTambahSukses(t *testing.T)  {
	result := tambah(20, 10)
	if result != 30 {
		t.Fail() // kalau gagal programnya akan dilanjut
	}
	fmt.Println("Test Selesai")
}

func TestTambahGagal(t *testing.T)  {
	result := tambah(20, 10)
	if result != 40 {
		t.FailNow() // kalau gagal testingnya berhenti
	}
	fmt.Println("Test Selesai")
}

func TestTambahError(t *testing.T)  {
	result := tambah(20, 10)
	if result != 40 {
		t.Error("Result harusnya 40")
	}
	fmt.Println("Test Selesai")
}

func TestTambahFatal(t *testing.T)  {
	result := tambah(20, 10)
	if result != 40 {
		t.Fatal("Result harusnya 40")
	}
	fmt.Println("Test Selesai")
}