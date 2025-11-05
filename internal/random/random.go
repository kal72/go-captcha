package random

import (
	"encoding/hex"
	"image/color"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Intn(i int) int {
	return r.Intn(i)
}

func Float64() float64 {
	return r.Float64()
}

// random nonce 8 character form byte
func Nonce() string {
	b := make([]byte, 4)
	for i := range b {
		b[i] = byte(r.Intn(256))
	}
	return hex.EncodeToString(b)
}

// random captcha code
func Code(n int) string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

// random strong color
func Color() color.RGBA {
	cases := r.Intn(6)
	switch cases {
	case 0:
		return color.RGBA{uint8(200 + r.Intn(55)), uint8(r.Intn(80)), uint8(r.Intn(80)), 255} // merah kuat
	case 1:
		return color.RGBA{uint8(r.Intn(80)), uint8(200 + r.Intn(55)), uint8(r.Intn(80)), 255} // hijau terang
	case 2:
		return color.RGBA{uint8(r.Intn(80)), uint8(r.Intn(80)), uint8(200 + r.Intn(55)), 255} // biru pekat
	case 3:
		return color.RGBA{uint8(180 + r.Intn(75)), uint8(r.Intn(60)), uint8(180 + r.Intn(75)), 255} // ungu
	case 4:
		return color.RGBA{uint8(220 + r.Intn(35)), uint8(140 + r.Intn(60)), uint8(r.Intn(50)), 255} // oranye tua
	default:
		return color.RGBA{uint8(160 + r.Intn(95)), uint8(r.Intn(120)), uint8(160 + r.Intn(95)), 255}
	}
}
