package test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestRead(t *testing.T)  {
	data, err := ioutil.ReadFile("D:/Data/data2.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func TestReadByWord(t *testing.T)  {
	file, err := os.Open("D:/Data/data2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestReadLineByLine(t *testing.T)  {
	file, err := os.Open("D:/Data/data2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
