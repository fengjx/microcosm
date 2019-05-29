package utils

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	hello := "你好，世界！ hello world"
	debyte := base64Encode([]byte(hello))
	fmt.Println(string(debyte))
	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}

func TestMd5(t *testing.T) {
	data := "admin"
	md5 := Md5(data)
	t.Log("md5:", md5)
	t.Log(Md5Check("admin", md5))
}
