package captcha

import "github.com/kal72/go-captcha/driver/memorystore"

type Config struct {
	SecretKey string
	Width     int     // width image
	Height    int     // height image
	Length    int     // length character text
	Expire    int     // in seconds
	FontSize  float64 // font size
	Store     Store   // cache
}

type Option func(*Config)

func defaultConfig(secret string) *Config {
	return &Config{
		SecretKey: secret,
		Width:     200,
		Height:    80,
		Length:    6,
		Expire:    60,
		FontSize:  38,
		Store:     memorystore.New(),
	}
}

// ===== Option functions =====

// Width and Height define the image size of the generated CAPTCHA in pixels.
func WithSize(width, height int) Option {
	return func(c *Config) {
		c.Width = width
		c.Height = height
	}
}

// Length defines the number of characters in the CAPTCHA code.
func WithLength(n int) Option {
	return func(c *Config) {
		c.Length = n
	}
}

// Expire the captcha code.
//
// Expire defines the CAPTCHA expiration time in seconds.
func WithExpire(seconds int) Option {
	return func(c *Config) {
		c.Expire = seconds
	}
}

// FontSize defines the font size used to render the CAPTCHA text in pixels.
func WithFontSize(fontSize float64) Option {
	return func(c *Config) {
		c.FontSize = fontSize
	}
}

// Available: memorystore or redisstore.
//
// To create a custom driver, implement the `Store` interface to define your own storage backend for CAPTCHA data.
func WithStore(store Store) Option {
	return func(c *Config) {
		c.Store = store
	}
}
