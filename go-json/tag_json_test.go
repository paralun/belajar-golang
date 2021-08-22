package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

func TestTagEncode(t *testing.T) {
	product := Product{
		Id:       "XYP01",
		Name:     "Aqua",
		Price:    3000,
		ImageUrl: "Http://paralun.com/aqua.png",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}

func TestTagDecode(t *testing.T) {
	jsonString := `{"id":"XYP01","name":"Aqua","price":3000,"image_url":"Http://paralun.com/aqua.png"}`
	jsonBytes := []byte(jsonString)
	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		return
	}

	fmt.Println(product)

}
