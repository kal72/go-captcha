package captcha

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/kal72/go-captcha/internal/assets"
	"github.com/kal72/go-captcha/internal/image"
	"github.com/kal72/go-captcha/internal/random"
	"github.com/kal72/go-captcha/internal/tokenutil"
)

type captcha struct {
	cfg *Config
}

// Defaut config
//
//	Width:     200
//	Height:    80
//	Length:    6
//	Expire:    60 //seconds
//	FontSize:  38
//	Store:     memorystore.New()
//
// Use New(secret string, opts ...Option) with option to manual config
func New(secret string, opts ...Option) *captcha {
	cfg := defaultConfig(secret)
	for _, opt := range opts {
		opt(cfg)
	}
	return &captcha{cfg: cfg}
}

func (c *captcha) Generate() (base64Image, text, token string, err error) {
	text = random.Code(c.cfg.Length)
	imgBytes, err := image.Draw(text, c.cfg.Width, c.cfg.Height, assets.DefaultFont, c.cfg.FontSize)
	if err != nil {
		return
	}
	base64Image = base64.StdEncoding.EncodeToString(imgBytes)
	base64Image = fmt.Sprintf("data:image/png;base64,%s", base64Image)

	exp := time.Now().Add(time.Duration(c.cfg.Expire) * time.Second).Unix()
	nonce := random.Nonce()
	data := tokenutil.Format(text, exp, nonce)
	token, err = tokenutil.Encrypt(c.cfg.SecretKey, []byte(data))
	if err != nil {
		return
	}

	return
}

func (c *captcha) Verify(answer string, token string) (err error) {
	dataDec, err := tokenutil.Decrypt(c.cfg.SecretKey, token)
	if err != nil {
		return
	}

	text, exp, _, err := tokenutil.ParseFormat(dataDec)
	if err != nil {
		return
	}

	if answer != text {
		err = errors.New("captcha invalid")
		return
	}

	if time.Now().Unix() > exp {
		err = errors.New("captcha expired")
		return
	}

	return
}
