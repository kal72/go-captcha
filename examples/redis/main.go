package main

import (
	"fmt"

	"github.com/kal72/go-captcha"
	"github.com/kal72/go-captcha/driver/redisstore"
)

func main() {
	cap := captcha.New("secret-key",
		captcha.WithStore(redisstore.New(redisstore.Config{
			Addr:     "localhost:6379",
			Password: "pass",
			DB:       0,
			Prefix:   "captcha",
		})),
	)
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
