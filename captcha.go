package captcha

import (
	"encoding/base64"
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

// Generate creates a new CAPTCHA image and returns its ID, text, and Base64-encoded image.
// It also stores the CAPTCHA code in the configured store (memory, Redis, etc.) until expiration.
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

// Verify checks whether the provided CAPTCHA code matches the one stored in the cache.
// If valid, the CAPTCHA entry will be removed from the store to prevent reuse.
func (c *captcha) Verify(answer string, token string) (valid bool, err error) {
	dataDec, err := tokenutil.Decrypt(c.cfg.SecretKey, token)
	if err != nil {
		return
	}

	text, exp, nonce, err := tokenutil.ParseFormat(dataDec)
	if err != nil {
		return
	}

	if answer != text {
		err = ErrCaptchaInvalid
		return
	}

	if time.Now().Unix() > exp {
		err = ErrCaptchaExpired
		return
	}

	value, _ := c.cfg.Store.Get(nonce)
	if value == token {
		err = ErrCaptchaClaimed
		return
	}

	ttl := time.Until(time.Unix(exp, 0))
	err = c.cfg.Store.Set(nonce, token, ttl)
	if err != nil {
		return
	}

	valid = true
	return
}
