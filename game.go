package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game interface {
}

type City struct {
	bg                                                       *ebiten.Image
	food, water, oxygen, power                               float32
	qualityOfLiving                                          float32
	icons, buttons                                           []*Button
	imgs                                                     []*ebiten.Image
	waterTime, foodTime, oxygenTime, powerTime, operatorTime time.Time
	operatorStatus                                           bool
}

func (c *City) Init() {
	c.LoadImgs()
	c.operatorStatus = true
	waterIcon := &Button{
		image: c.imgs[0],
		x:     50,
		y:     20,
	}
	foodIcon := &Button{
		image: c.imgs[1],
		x:     150,
		y:     20,
	}
	oxygenIcon := &Button{
		image: c.imgs[2],
		x:     250,
		y:     20,
	}
	powerIcon := &Button{
		image: c.imgs[3],
		x:     350,
		y:     20,
	}

	operatorIcon := &Button{
		image:   c.imgs[12],
		x:       450,
		y:       20,
		onClick: func() { c.operatorStatus = true },
	}

	c.waterTime = time.Now()
	c.foodTime = time.Now()
	c.oxygenTime = time.Now()
	c.powerTime = time.Now()
	c.operatorTime = time.Now()

	c.icons = append(c.icons, waterIcon, foodIcon, oxygenIcon, powerIcon, operatorIcon)

	logButton := &Button{
		image: c.imgs[13],
		x:     (WIDTH - c.imgs[13].Bounds().Dx()) / 2,
		y:     HEIGHT - 100,
		text:  "Logs & reports",
	}

	c.buttons = append(c.buttons, logButton)
}

func (c *City) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(c.bg, 2, 2)
	screen.DrawImage(c.bg, opBG)

	for i, icon := range c.icons {
		if i == len(c.icons)-1 {
			break
		}
		icon.Draw(screen)

	}

	for _, button := range c.buttons {
		button.Draw(screen)
	}

	if !c.operatorStatus {
		c.icons[4].Draw(screen)
	}
}

func (c *City) ManageResources() {

}

func (c *City) MonitorResources() {
	if c.water >= 75 {
		c.icons[0].image = c.imgs[0]
	} else if c.water >= 40 {
		c.icons[0].image = c.imgs[4]
	} else {
		c.icons[0].image = c.imgs[8]
	}

	if c.food >= 75 {
		c.icons[1].image = c.imgs[1]
	} else if c.food >= 25 {
		c.icons[1].image = c.imgs[5]
	} else {
		c.icons[1].image = c.imgs[9]
	}

	if c.oxygen >= 75 {
		c.icons[2].image = c.imgs[2]
	} else if c.oxygen >= 50 {
		c.icons[2].image = c.imgs[6]
	} else {
		c.icons[2].image = c.imgs[10]
	}

	if c.power >= 75 {
		c.icons[3].image = c.imgs[3]
	} else if c.power >= 50 {
		c.icons[3].image = c.imgs[7]
	} else {
		c.icons[3].image = c.imgs[11]
	}
}

func (c *City) UpdateResources() {
	if time.Since(c.waterTime) >= 10*time.Second {
		c.water = c.water - 5
		c.KeepInBounds(c.water)
		c.waterTime = time.Now()
	}

	if time.Since(c.foodTime) >= 20*time.Second {
		c.food = c.food - 5
		c.KeepInBounds(c.food)
		c.foodTime = time.Now()
	}

	if time.Since(c.oxygenTime) >= 6*time.Second {
		c.oxygen = c.oxygen - 5
		c.KeepInBounds(c.oxygen)
		c.oxygenTime = time.Now()
	}

	if time.Since(c.powerTime) >= 20*time.Second {
		c.power = c.power - 5
		c.KeepInBounds(c.power)
		c.powerTime = time.Now()
	}
	c.MonitorResources()
	c.CheckOperator()
}

func (c *City) KeepInBounds(resource float32) {
	if resource < 0 {
		resource = 0
	} else if resource > 100 {
		resource = 100
	}
}

func (c *City) CheckOperator() {
	if time.Since(c.operatorTime) >= 30*time.Second {
		c.operatorStatus = false
		c.operatorTime = time.Now()
	}
	c.icons[4].Update()
}

func (c *City) ShowDashboards() {

}

func (c *City) ShowLogs() {

}

func (c *City) LoadImgs() {
	// good status of resources icons
	waterImgG, _, _ := ebitenutil.NewImageFromFile(ICON1)
	foodImgG, _, _ := ebitenutil.NewImageFromFile(ICON2)
	oxygenImgG, _, _ := ebitenutil.NewImageFromFile(ICON3)
	powerImgG, _, _ := ebitenutil.NewImageFromFile(ICON4)

	// ok status of resources icons
	waterImgO, _, _ := ebitenutil.NewImageFromFile(ICON5)
	foodImgO, _, _ := ebitenutil.NewImageFromFile(ICON6)
	oxygenImgO, _, _ := ebitenutil.NewImageFromFile(ICON7)
	powerImgO, _, _ := ebitenutil.NewImageFromFile(ICON8)

	// bad status of resources icons
	waterImgB, _, _ := ebitenutil.NewImageFromFile(ICON9)
	foodImgB, _, _ := ebitenutil.NewImageFromFile(ICON10)
	oxygenImgB, _, _ := ebitenutil.NewImageFromFile(ICON11)
	powerImgB, _, _ := ebitenutil.NewImageFromFile(ICON12)

	// appears everytime oxygen changes status
	operatorImg, _, _ := ebitenutil.NewImageFromFile(ICON13)

	// buttons
	logImg, _, _ := ebitenutil.NewImageFromFile(BUTTON4)

	c.imgs = append(c.imgs,
		waterImgG, foodImgG, oxygenImgG, powerImgG,
		waterImgO, foodImgO, oxygenImgO, powerImgO,
		waterImgB, foodImgB, oxygenImgB, powerImgB,
		operatorImg, logImg)
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
