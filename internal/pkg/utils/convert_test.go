package utils

import "testing"

func TestConvert(t *testing.T) {
	var a int64 = 1234
	t.Log(convInt64ToBytes(a))
}
