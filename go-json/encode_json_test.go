package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
func TestEncode(t *testing.T) {
	logJson("James Kusmambang")
	logJson(10)
	logJson(true)
	logJson([]string{"Ini", "Bapa", "Budi"})
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
	Hobbies   []string
}

func TestEncodeObject(t *testing.T) {
	custome := Customer{
		FirstName: "Indah",
		LastName:  "Pantai Kapuk",
		Age:       23,
		Married:   false,
		Hobbies:   []string{"Makan", "Nonton", "Tidur"},
	}

	bytes, err := json.Marshal(custome)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
