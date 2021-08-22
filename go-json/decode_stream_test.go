package go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecodeStream(t *testing.T) {
	reader, err := os.Open("product.json")
	if err != nil {
		return
	}

	decoder := json.NewDecoder(reader)
	product := &Product{}
	decoder.Decode(product)
	fmt.Println(product)
}
