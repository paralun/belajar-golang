package test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMkdir(t *testing.T)  {
	//err := os.Mkdir("/Data", 0755)
	err := os.Mkdir("D:/Data", 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMkdirAll(t *testing.T)  {
	//err := os.MkdirAll("/Data/Dir", 0755)
	err := os.MkdirAll("D:/Data/Dir1", 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCD(t *testing.T)  {
	os.Chdir("D:/Data/Dir1")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)
}
