package main

import (
	"image"
	"os"
	"strings"

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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(image, op)
}

func ResizeImg(path string, w, h uint) *ebiten.Image {
	openedFile, _ := os.Open(path)
	img, _, _ := image.Decode(openedFile)
	resizedImg := resize.Resize(w, h, img, resize.Lanczos3)
	ebitenImg := ebiten.NewImageFromImage(resizedImg)
	return ebitenImg
}

func PlaceImg(x, y int, img, screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, options)
}

func WrapText(text string, width int) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return lines
	}

	currentLine := ""
	currentLineLength := 0

	for _, word := range words {
		if currentLineLength+len(word) > width {
			lines = append(lines, currentLine)
			currentLine = word
			currentLineLength = len(word)
		} else {
			if currentLineLength > 0 {
				currentLine += " "
				currentLineLength++
			}
			currentLine += word
			currentLineLength += len(word)
		}
	}

	if currentLineLength > 0 {
		lines = append(lines, currentLine)
	}

	return lines
}
