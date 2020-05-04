package utility

import (
	"math/rand"
	"time"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func GenerateColor() string {
	rand.Seed(time.Now().UnixNano())
	c := colorful.Hsv(rand.Float64()*360.0, 0.4, 0.8)
	return c.Hex()
}
