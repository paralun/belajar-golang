package go_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed tabel.png
var gambar []byte

//go:embed files/satu.txt
//go:embed files/dua.txt
//go:embed files/tiga.txt
var files embed.FS

//go:embed files/*.txt
var path embed.FS

func TestEmbedString(t *testing.T)  {
	fmt.Println(version)
}

func TestEmbedByte(t *testing.T)  {
	err := ioutil.WriteFile("new_tabel.png", gambar, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestEmbedMultipleFile(t *testing.T)  {
	satu, _ := files.ReadFile("files/satu.txt")
	fmt.Println(string(satu))

	dua, _ := files.ReadFile("files/dua.txt")
	fmt.Println(string(dua))

	tiga, _ := files.ReadFile("files/tiga.txt")
	fmt.Println(string(tiga))
}

func TestEmbedPathMatch(t *testing.T)  {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			f, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(f))
		}
	}
}