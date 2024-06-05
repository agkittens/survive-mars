package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game interface {
}

type City struct {
	bg                                                       *ebiten.Image
	food, water, oxygen, power                               float32
	qualityOfLiving                                          float32
	icons, buttons, windows                                  []*Button
	imgs                                                     []*ebiten.Image
	waterTime, foodTime, oxygenTime, powerTime, operatorTime time.Time
	operatorStatus                                           bool
	logs                                                     []string
	isLogShowed, isDashbShowed                               bool
}

func (c *City) Init() {
	c.LoadImgs()
	c.CreateButtons()
	c.operatorStatus = true
	c.isDashbShowed = false
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

	for i := 0; i < len(c.buttons); i++ {
		c.buttons[i].Draw(screen)
	}

	if !c.operatorStatus {
		c.icons[4].Draw(screen)
	}

	c.ShowLogs(screen)
	c.ShowDashboards(screen)
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
		c.UpdateLogs("water", "down")
	}

	if time.Since(c.foodTime) >= 20*time.Second {
		c.food = c.food - 5
		c.KeepInBounds(c.food)
		c.foodTime = time.Now()
		c.UpdateLogs("food", "down")

	}

	if time.Since(c.oxygenTime) >= 1*time.Second {
		c.oxygen = c.oxygen - 5
		c.KeepInBounds(c.oxygen)
		c.oxygenTime = time.Now()
		c.UpdateLogs("oxygen", "down")

	}

	if time.Since(c.powerTime) >= 20*time.Second {
		c.power = c.power - 5
		c.KeepInBounds(c.power)
		c.powerTime = time.Now()
		c.UpdateLogs("power", "down")

	}
	for i := 0; i < len(c.buttons); i++ {
		c.buttons[i].Update()
	}
	for i, icon := range c.icons {
		if i == len(c.icons)-1 {
			break
		}
		icon.Update()
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
		c.logs = append(c.logs, "check on operator")
	}
	c.icons[4].Update()
}

func (c *City) ShowDashboards(screen *ebiten.Image) {
	log.Print(c.isDashbShowed)
	if c.isDashbShowed {
		opBG := AdjustSize(c.bg, 2, 2)
		screen.DrawImage(c.bg, opBG)
		c.windows[1].Draw(screen)
		c.windows[2].Draw(screen)
		c.windows[3].Draw(screen)

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			c.isDashbShowed = false

		}
	}

}

func (c *City) ShowLogs(screen *ebiten.Image) {
	if c.isLogShowed {
		c.buttons[0].image = c.imgs[14]
		c.windows[0].Draw(screen)
		logWindowHeight := c.imgs[15].Bounds().Dy() - 32

		for i, text := range c.logs {
			y := c.windows[0].y + (i+1)*16 - 16
			if y >= c.windows[0].y && y <= c.windows[0].y+logWindowHeight-32 {
				DisplayText(WIDTH/2-len(text)*4, y+32, 16, text, screen, color.White)
			} else {
				c.logs = c.logs[i+1:]
				break
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			c.isLogShowed = false
			c.buttons[0].image = c.imgs[13]

		}
	}
}

func (c *City) UpdateLogs(resource, status string) {
	if status == "up" {
		c.logs = append(c.logs, string(resource)+" was added")
	} else if status == "down" {
		c.logs = append(c.logs, string(resource)+" decreased")
	}
}

func (c *City) CreateButtons() {
	waterIcon := &Button{
		image:   c.imgs[0],
		x:       50,
		y:       20,
		onClick: func() { c.isDashbShowed = true },
	}
	foodIcon := &Button{
		image:   c.imgs[1],
		x:       150,
		y:       20,
		onClick: func() { c.isDashbShowed = true },
	}
	oxygenIcon := &Button{
		image:   c.imgs[2],
		x:       250,
		y:       20,
		onClick: func() { c.isDashbShowed = true },
	}
	powerIcon := &Button{
		image:   c.imgs[3],
		x:       350,
		y:       20,
		onClick: func() { c.isDashbShowed = true },
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
		image:   c.imgs[13],
		x:       (WIDTH - c.imgs[13].Bounds().Dx()) / 2,
		y:       HEIGHT - 100,
		text:    "Logs & reports",
		onClick: func() { c.isLogShowed = true },
	}

	logWindow := &Button{
		image: c.imgs[15],
		x:     (WIDTH - c.imgs[15].Bounds().Dx()) / 2,
		y:     HEIGHT / 4,
	}

	dashboard1 := &Button{
		image: c.imgs[17],
		x:     (WIDTH - c.imgs[16].Bounds().Dx()) / 6,
		y:     0,
	}

	dashboard2 := &Button{
		image: c.imgs[16],
		x:     (WIDTH/2 + 10),
		y:     0,
	}

	dashboard3 := &Button{
		image: c.imgs[16],
		x:     (WIDTH/2 + 10),
		y:     c.imgs[16].Bounds().Dy(),
	}

	c.buttons = append(c.buttons, logButton)
	c.windows = append(c.windows, logWindow, dashboard1, dashboard2, dashboard3)
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
	openedFile, _ := os.Open(BUTTON4)
	img, _, _ := image.Decode(openedFile)
	logImg := ResizeImg(img, 260, 52)

	openedFile, _ = os.Open(BUTTON5)
	img, _, _ = image.Decode(openedFile)
	logImg2 := ResizeImg(img, 260, 52)

	// log window
	logWindowImg, _, _ := ebitenutil.NewImageFromFile(LOG_WINDOW)

	// dashboards
	dashboardImg1, _, _ := ebitenutil.NewImageFromFile(DASHBOARD1)

	openedFile, _ = os.Open(DASHBOARD2)
	img, _, _ = image.Decode(openedFile)
	dashboardImg2 := ResizeImg(img, 405, 750)

	c.imgs = append(c.imgs,
		waterImgG, foodImgG, oxygenImgG, powerImgG,
		waterImgO, foodImgO, oxygenImgO, powerImgO,
		waterImgB, foodImgB, oxygenImgB, powerImgB,
		operatorImg, logImg, logImg2, logWindowImg,
		dashboardImg1, dashboardImg2)
}
