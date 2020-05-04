package utility

import (
	"math/rand"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func GenerateColor() string {
	c := colorful.Hsv(rand.Float64()*360.0, 0.8, 0.8)
	return c.Hex()
}
