package utils

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

// 生成二维码文件
func GenerateQRCodeFile(content, fileName string, size int) {
	err := qrcode.WriteFile(content, qrcode.Medium, size, fileName)
	if err != nil {
		log.Fatal(err)
	}
}

// 生成带颜色二维码
func GenerateQRCodeColorFile(content, fileName string, size int) {
	err := qrcode.WriteColorFile(content, qrcode.Medium, size, color.White, color.RGBA{
		R: 126,
		G: 12,
		B: 245,
		A: 200,
	}, fileName)
	if err != nil {
		log.Fatal(err)
	}
}

// 生成二维码字节流
func GenerateQRCodeBytes(content string, size int) ([]byte, error) {
	bytes, err := qrcode.Encode(content, qrcode.Medium, size)
	if err != nil {
		log.Fatal(err)
	}
	return bytes, nil
}

// 生成 base64字符串
func GenerateQRCodeBase64(content string, size int) string {
	bytes, _ := GenerateQRCodeBytes(content, size)
	encodingString := base64.StdEncoding.EncodeToString(bytes)
	return encodingString
}
