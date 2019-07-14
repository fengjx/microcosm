package protoc

import (
	"encoding/json"
	"testing"
)

func TestPage(t *testing.T) {
	var list = []string{"a", "b", "c"}
	page := BuildPage(100, 10, 2, list)
	t.Log(json.Marshal(page))
}
