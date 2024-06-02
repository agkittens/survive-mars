package main

import (
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Window struct {
	background *ebiten.Image
}

func (w *Window) Init() {
	w.background, _, _ = ebitenutil.NewImageFromFile(BG)
}

func (w *Window) Update() error {
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(w.background, 2, 2)
	screen.DrawImage(w.background, opBG)
}
func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}
