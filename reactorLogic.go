package main

import "github.com/hajimehoshi/ebiten/v2"

type Reactor struct {
	bg *ebiten.Image
}

func (r *Reactor) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(r.bg, 2, 2)
	screen.DrawImage(r.bg, opBG)
}
