package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nfnt/resize"
)

func AdjustSize(img *ebiten.Image, divX int, divY int) *ebiten.DrawImageOptions {
	size := img.Bounds().Size()
	posX := (WIDTH - size.X) / divX
	posY := (HEIGHT - size.Y) / divY
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posX), float64(posY))
	return op
}

func CreateRect(x, y int, scaleX, scaleY float64, screen, image *ebiten.Image) {
	// img, _, _ := ebitenutil.NewImageFromFile(image)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(image, op)
}

func ResizeImg(img image.Image, w, h uint) *ebiten.Image {
	resizedImg := resize.Resize(w, h, img, resize.Lanczos3)
	ebitenImg := ebiten.NewImageFromImage(resizedImg)
	return ebitenImg
}
