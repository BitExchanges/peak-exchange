package utils

import (
	"errors"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"math/rand"
	"strings"
	"time"
)

var store = base64Captcha.DefaultMemStore

var captchaChinese = base64Captcha.DriverChinese{
	Height:          40,
	Width:           155,
	NoiseCount:      30,
	ShowLineOptions: 0,
	Length:          2,
	Source:          "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,不想要,的值",
	BgColor: &color.RGBA{
		R: 125,
		G: 125,
		B: 0,
		A: 118,
	},
	Fonts: []string{"wqy-microhei.ttc"},
}

var captchaDigit = base64Captcha.DriverDigit{
	Height:   40,
	Width:    120,
	Length:   4,
	MaxSkew:  0.2,
	DotCount: 76,
}

var captchaString = base64Captcha.DriverString{
	Height:          45,
	Width:           105,
	NoiseCount:      20,
	ShowLineOptions: 0,
	Length:          6,
	Source:          "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	BgColor: &color.RGBA{
		R: 52,
		G: 97,
		B: 41,
		A: 0,
	},
	Fonts: nil,
}

func GenerateCaptcha(typ string) (string, string, error) {
	var driver base64Captcha.Driver
	switch typ {
	case "string":
		driver = captchaString.ConvertFonts()
	case "digit":
		driver = &captchaDigit
	case "chinese":
		driver = captchaChinese.ConvertFonts()
	}

	c := base64Captcha.NewCaptcha(driver, store)

	id, data, err := c.Generate()
	if err != nil {
		return "", "", errors.New("生成验证码失败")
	}
	return id, data, nil
}

// verify base64 captcha code
func VerifyCaptcha(id, value string) error {
	if !store.Verify(id, value, true) {
		return errors.New("验证码错误")
	}
	return nil
}

//生成任意位纯数字验证码
func GenerateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
