package image

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"

	"github.com/golang/freetype"
	"github.com/kal72/go-captcha/internal/random"
)

// line
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, col color.RGBA, thickness int, width, height int) {
	dx := math.Abs(float64(x2 - x1))
	dy := math.Abs(float64(y2 - y1))
	sx, sy := 1, 1
	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy
	for {
		for t := -thickness; t <= thickness; t++ {
			if x1+t >= 0 && x1+t < width && y1+t >= 0 && y1+t < height {
				img.Set(x1+t, y1, col)
				img.Set(x1, y1+t, col)
			}
		}
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func drawNoise(img *image.RGBA, width, height int) {
	img.Set(random.Intn(width), random.Intn(height), random.Color())
}

// draw captcha image
func Draw(text string, width, height int, font []byte, fontSize float64) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bgColor := color.RGBA{240, 240, 240, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	for i := 0; i < 200; i++ {
		drawNoise(img, width, height)
	}

	for i := 0; i < 5; i++ {
		drawLine(img, random.Intn(width), random.Intn(height),
			random.Intn(width), random.Intn(height),
			random.Color(), 1, width, height)
	}

	f, err := freetype.ParseFont(font)
	if err != nil {
		return nil, err
	}

	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFont(f)
	ctx.SetFontSize(fontSize)
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)

	textWidth := float64(width) * 0.85
	charSpacing := textWidth / float64(len(text))
	startX := (float64(width) - textWidth) / 2
	yBase := float64(height) * 0.7

	for i, c := range text {
		ctx.SetSrc(image.NewUniform(random.Color()))

		randomSpacing := charSpacing * (0.8 + random.Float64()*0.4)
		if i > 0 {
			startX += randomSpacing
		}

		waveOffset := math.Sin(float64(i)*random.Float64()*math.Pi*0.5) * 6
		jitterY := random.Float64()*4 - 2
		y := yBase + waveOffset + jitterY

		pt := freetype.Pt(int(startX), int(y))
		_, _ = ctx.DrawString(string(c), pt)
	}

	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes(), nil
}
