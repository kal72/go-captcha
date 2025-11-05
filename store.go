package captcha

import "time"

// Store defines the interface for CAPTCHA data storage.
// To create a custom driver, implement this interface.
type Store interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}
