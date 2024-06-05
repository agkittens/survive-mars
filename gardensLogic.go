package main

import "github.com/hajimehoshi/ebiten/v2"

type Gardens struct {
	bg *ebiten.Image
}

func (g *Gardens) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(g.bg, 2, 2)
	screen.DrawImage(g.bg, opBG)
}
