package utils

import (
	"bytes"
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func InitDraw() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.ArcadeN_ttf))
	if err != nil {
		panic(err)
	}
	mplusFaceSource = s
}

func DrawText(screen *ebiten.Image, x float64, y float64, fontSize float64, value string, textColor color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(textColor)
	text.Draw(screen, value, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   fontSize,
	}, op)
}
