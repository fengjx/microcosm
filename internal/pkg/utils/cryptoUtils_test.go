package utils

import (
	"testing"
)

func TestBase64(t *testing.T) {
	hello := "你好，世界！ hello world"
	debyte := base64Encode([]byte(hello))
	t.Log(string(debyte))
	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		t.Error(err.Error())
	}

	if hello != string(enbyte) {
		t.Log("hello is not equal to enbyte")
	}

	t.Log(string(enbyte))
}

func TestMd5(t *testing.T) {
	data := "30070831ffda3272-fds2-432k-482e-366d3425a942"
	md5 := Md5(data)
	t.Log("md5:", md5)
	t.Log(Md5Check("admin", md5))
}
