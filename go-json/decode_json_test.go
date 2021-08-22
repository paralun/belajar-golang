package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonString := `{"FirstName":"Indah","LastName":"Pantai Kapuk","Age":23,"Married":false,"Hobbies":["Makan","Nonton","Tidur"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		return
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}
