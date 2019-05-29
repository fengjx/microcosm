package utils

import "testing"

func TestRandString(t *testing.T) {
	r1 := RandString(3)
	r2 := RandString(3)
	t.Log(r1, r2)
}
