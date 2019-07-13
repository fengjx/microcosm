package utils

import (
	"testing"
	"time"
)

func TestJwtToken(t *testing.T) {

	token, _ := GenJwtToken("fengjx", "1234")
	t.Log(token)
	time.Sleep(3 * time.Second)
	claims, _ := ParseJwtToken(token)
	t.Log(claims)

}
