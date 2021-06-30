package test

import (
	"log"
	"os"
	"testing"
)

func TestRemoveFile(t *testing.T)  {
	err := os.Remove("D:/Data/data1.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveDirEmpty(t *testing.T)  {
	err := os.Remove("D:/Data/Dir2")
	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveDirAll(t *testing.T)  {
	err := os.RemoveAll("D:/Data/Dir1")
	if err != nil {
		log.Fatal(err)
	}
}
