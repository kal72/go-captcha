package main

import (
	"fmt"

	"github.com/kal72/go-captcha"
)

func main() {
	cap := captcha.New("qwertyasdfzxcv1234")
	base64Image, text, token, err := cap.Generate()
	if err != nil {
		panic(err)
	}

	fmt.Println("text: " + text)
	fmt.Println("token: " + token)
	fmt.Println("image: " + base64Image)

	_, err = cap.Verify(text, token)
	if err != nil {
		panic("verify failed: " + err.Error())
	}

	fmt.Println("verify: success")
}
