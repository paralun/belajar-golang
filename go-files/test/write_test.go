package test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestWrite(t *testing.T)  {
	//file, err := os.Create("./data1.txt")
	file, err := os.Create("D:/Data/data1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//file.Write([]byte("Create to File"))
	n, err := file.WriteString("Create to File with String")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
}

func TestWriteIOUtil(t *testing.T)  {
	text := []byte("Create file with IOUtil")
	err := ioutil.WriteFile("D:/Data/data2.txt", text, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func TestWriteBuffer(t *testing.T)  {
	file, err := os.Create("./data1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("Create data ke file")
	if err != nil {
		log.Fatal(err)
	}

	defer writer.Flush()
}
