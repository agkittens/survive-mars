package main

import (
	"bytes"
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	image   *ebiten.Image
	x, y    int
	text    string
	onClick func()
}

func (b *Button) Update() {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()
		if (posX >= b.x && posX <= b.x+b.image.Bounds().Dx()) && (posY >= b.y && posY <= b.y+b.image.Bounds().Dy()) {
			b.onClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	CreateRect(b.x, b.y, 1, 1, screen, b.image)
	textX := WIDTH/2 - len(b.text)*6
	textY := b.y + b.image.Bounds().Dy()/3
	DisplayText(textX, textY, 24, b.text, screen, color.White)
}

func DisplayText(x, y, size int, msg string, screen *ebiten.Image, color color.Color) {
	mplusFaceSource, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(size),
	}, op)
}
