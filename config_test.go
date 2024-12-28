package ic_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aviate-labs/ic-go"
)

func TestConfig(t *testing.T) {
	raw, err := os.ReadFile("./config.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(raw, new(ic.Config)); err != nil {
		t.Error(err)
	}
}
