package main

import "github.com/hajimehoshi/ebiten/v2"

type Transportation struct {
	bg *ebiten.Image
}

func (t *Transportation) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(t.bg, 2, 2)
	screen.DrawImage(t.bg, opBG)
}
