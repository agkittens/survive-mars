package main

import "github.com/hajimehoshi/ebiten/v2"

type Game interface {
}

type City struct {
	bg *ebiten.Image
}

func (c *City) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(c.bg, 2, 2)
	screen.DrawImage(c.bg, opBG)
}

type Gardens struct {
	bg *ebiten.Image
}

func (g *Gardens) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(g.bg, 2, 2)
	screen.DrawImage(g.bg, opBG)
}

type Reactor struct {
	bg *ebiten.Image
}

func (r *Reactor) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(r.bg, 2, 2)
	screen.DrawImage(r.bg, opBG)
}

type Transportation struct {
	bg *ebiten.Image
}

func (t *Transportation) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(t.bg, 2, 2)
	screen.DrawImage(t.bg, opBG)
}
