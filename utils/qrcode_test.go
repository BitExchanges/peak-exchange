package utils

import "testing"

func TestGenerateQRCodeFile(t *testing.T) {
	GenerateQRCodeFile("http://www.baidu.com", "baidu.png", 256)
}

func TestGenerateQRCodeColorFile(t *testing.T) {
	GenerateQRCodeColorFile("http://www.naidu.com", "color.png", 256)
}

func TestGenerateQRCodeBase64(t *testing.T) {
	str := GenerateQRCodeBase64("http://www.baidu.com", 256)
	t.Log(str)
}
