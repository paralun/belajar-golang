package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncodeStream(t *testing.T) {
	create, err := os.Create("output.json")
	if err != nil {
		return
	}
	encoder := json.NewEncoder(create)
	product := Product{
		Id:       "XZP02",
		Name:     "Kopi Hitam",
		Price:    1500,
		ImageUrl: "http://paralun.com/kopi.png",
	}
	_ = encoder.Encode(product)
}
