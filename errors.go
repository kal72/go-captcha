package captcha

import "errors"

var (
	ErrCaptchaNotFound = errors.New("captcha not found")
	ErrCaptchaExpired  = errors.New("captcha expired")
	ErrCaptchaInvalid  = errors.New("captcha invalid")
	ErrCaptchaClaimed  = errors.New("captcha already used")
)
