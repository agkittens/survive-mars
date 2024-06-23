package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Gardens struct {
	bg   *ebiten.Image
	imgs []*ebiten.Image
	//icons, buttons, windows []*Button
	buttons, windows, icons []*Button
	plantsButtons           []*Button
}

func (g *Gardens) Init() {
	g.LoadImgs()
	g.CreateButtons()
}

func (g *Gardens) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(g.bg, 2, 2)
	screen.DrawImage(g.bg, opBG)

	for i := 0; i < len(g.buttons); i++ {
		g.buttons[i].Draw(screen)
	}
	for i := 0; i < len(g.windows); i++ {
		g.windows[i].Draw(screen)
	}
	for i := 0; i < len(g.plantsButtons); i++ {
		g.plantsButtons[i].Draw(screen)
	}
}

func (g *Gardens) Update() {
	for i := 0; i < len(g.buttons); i++ {
		g.buttons[i].Update()
	}
}

func (g *Gardens) CreateButtons() {

	plantIcon0 := &Button{
		image: g.imgs[4],
		x:     185,
		y:     170,
	}

	plantIcon1 := &Button{
		image: g.imgs[4],
		x:     290,
		y:     170,
	}
	plantIcon2 := &Button{
		image: g.imgs[4],
		x:     385,
		y:     170,
	}
	g.plantsButtons = append(g.plantsButtons, plantIcon0, plantIcon1, plantIcon2)

	graphsButton := &Button{
		image: g.imgs[0],
		x:     3 * (WIDTH - g.imgs[0].Bounds().Dx()) / 4,
		y:     HEIGHT - 100,
		text:  "Show graphs",
	}
	g.buttons = append(g.buttons, graphsButton)

	plantsWindow := &Button{
		image: g.imgs[2],
		x:     (WIDTH - g.imgs[2].Bounds().Dx()) / 6,
		y:     25,
	}

	plantsImg := &Button{
		image: g.imgs[3],
		x:     55 + (WIDTH-g.imgs[2].Bounds().Dx())/6,
		y:     50,
	}

	g.windows = append(g.windows, plantsWindow, plantsImg)
}

func (g *Gardens) LoadImgs() {
	logImg := ResizeImg(BUTTON4, 260, 52)
	logImg2 := ResizeImg(BUTTON5, 260, 52)

	plantsWindowImg := ResizeImg(DASHBOARD2, 405, 700)
	plantsImg := ResizeImg(PLANTS, 300, 650)

	plantIcon, _, _ := ebitenutil.NewImageFromFile(ICON15)
	g.imgs = append(g.imgs, logImg, logImg2, plantsWindowImg, plantsImg, plantIcon)

}
