package main

import (
	"fmt"
	"image/color"
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
	resources                                                []float32
	qualityOfLiving                                          float32
	icons, buttons, windows                                  []*Button
	imgs                                                     []*ebiten.Image
	waterTime, foodTime, oxygenTime, powerTime, operatorTime time.Time
	operatorStatus                                           bool
	logs                                                     []string
	isLogShowed, isDashbShowed                               bool
	currentResource                                          string
	currentResourceIdx                                       int
	resourceDescription                                      []string
	resourceDescriptionIdx                                   int
}

func (c *City) Init() {
	c.LoadImgs()
	c.CreateButtons()
	c.GetResourceDescription()
	c.operatorStatus = true
	c.isDashbShowed = false
	c.qualityOfLiving = 100
	c.resources = append(c.resources, c.water, c.food, c.oxygen, c.power)
	c.resourceDescriptionIdx = 0
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

	c.buttons[0].Draw(screen)

	if !c.operatorStatus {
		c.icons[5].Draw(screen)
	}

	c.ShowLogs(screen)
	c.ShowDashboards(screen)
}

func (c *City) ManageResources() {

}

func (c *City) MonitorResources() {
	if c.resources[0] >= 75 {
		c.icons[0].image = c.imgs[0]
	} else if c.resources[0] >= 40 {
		c.icons[0].image = c.imgs[4]
	} else {
		c.icons[0].image = c.imgs[8]
	}

	if c.resources[1] >= 75 {
		c.icons[1].image = c.imgs[1]
	} else if c.resources[1] >= 25 {
		c.icons[1].image = c.imgs[5]
	} else {
		c.icons[1].image = c.imgs[9]
	}

	if c.resources[2] >= 75 {
		c.icons[2].image = c.imgs[2]
	} else if c.resources[2] >= 50 {
		c.icons[2].image = c.imgs[6]
	} else {
		c.icons[2].image = c.imgs[10]
	}

	if c.resources[3] >= 75 {
		c.icons[3].image = c.imgs[3]
	} else if c.resources[3] >= 50 {
		c.icons[3].image = c.imgs[7]
	} else {
		c.icons[3].image = c.imgs[11]
	}
}

func (c *City) UpdateResources() {
	if time.Since(c.waterTime) >= 10*time.Second {
		c.resources[0] = c.resources[0] - 5
		c.KeepInBounds(0)
		c.waterTime = time.Now()
		c.UpdateLogs("water", "down")

	}

	if time.Since(c.foodTime) >= 20*time.Second {
		c.resources[1] = c.resources[1] - 5
		c.KeepInBounds(1)
		c.foodTime = time.Now()
		c.UpdateLogs("food", "down")

	}

	if time.Since(c.oxygenTime) >= 6*time.Second {
		c.resources[2] = c.resources[2] - 5
		c.KeepInBounds(2)
		c.oxygenTime = time.Now()
		c.UpdateLogs("oxygen", "down")

	}

	if time.Since(c.powerTime) >= 20*time.Second {
		c.resources[3] = c.resources[3] - 5
		c.KeepInBounds(3)
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

func (c *City) KeepInBounds(resourceIdx int) {
	if c.resources[resourceIdx] <= 0 {
		c.resources[resourceIdx] = 0
	} else if c.resources[resourceIdx] >= 100 {
		c.resources[resourceIdx] = 100
	}
}

func (c *City) CheckOperator() {
	if time.Since(c.operatorTime) >= 30*time.Second && c.operatorStatus {
		c.operatorStatus = false
		c.operatorTime = time.Now()
		c.logs = append(c.logs, "check on operator")
	}
	//if time.Since(c.operatorTime) >= 30*time.Second && !c.operatorStatus {
	//	currentState = StateExit

	//}

	c.icons[5].Update()
}

func (c *City) ShowDashboards(screen *ebiten.Image) {
	if c.isDashbShowed {
		opBG := AdjustSize(c.bg, 2, 2)
		screen.DrawImage(c.bg, opBG)
		c.windows[1].Draw(screen)
		c.windows[2].Draw(screen)
		c.windows[3].Draw(screen)

		c.buttons[1].Draw(screen)
		c.buttons[2].Draw(screen)
		c.buttons[3].Draw(screen)

		DisplayText(c.windows[1].x+c.windows[1].image.Bounds().Dx()/2-len(c.currentResource)*9, c.windows[1].y+60, 36, c.currentResource, screen, color.White)
		msg := fmt.Sprintf("%.2f", c.resources[c.resourceDescriptionIdx])
		DisplayText(c.windows[2].x+c.windows[2].image.Bounds().Dx()/2-len(msg)*9, c.windows[2].y+60, 36, msg, screen, color.White)

		DisplayText(c.windows[2].x+c.windows[2].image.Bounds().Dx()/2-len("Quality of living")*8, c.windows[2].y+150, 36, "Quality of living", screen, color.White)
		msg = fmt.Sprintf("%.2f", c.qualityOfLiving)
		DisplayText(c.windows[2].x+c.windows[2].image.Bounds().Dx()/2-len(msg)*9, c.windows[2].y+200, 36, msg, screen, color.White)

		xPos := c.windows[1].x + 5 + (c.imgs[16].Bounds().Dx()-c.imgs[c.currentResourceIdx].Bounds().Dx())/2
		PlaceImg(xPos, 370, c.imgs[c.currentResourceIdx], screen)

		text := WrapText(c.resourceDescription[c.resourceDescriptionIdx], 35)
		for i, word := range text {
			y := c.windows[1].y + 100 + (i+1)*16
			DisplayText(xPos, y+32, 16, word, screen, color.White)
		}

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
		c.logs = append(c.logs, resource+" was added")
	} else if status == "down" {
		c.logs = append(c.logs, resource+" decreased")
	}
}

func (c *City) CreateButtons() {

	// city main screen related
	waterIcon := &Button{
		image: c.imgs[0],
		x:     50,
		y:     20,
		onClick: func() {
			c.isDashbShowed = true
			c.currentResource = "water"
			c.currentResourceIdx = 18
			c.resourceDescriptionIdx = 0
		},
	}
	foodIcon := &Button{
		image: c.imgs[1],
		x:     150,
		y:     20,
		onClick: func() {
			c.isDashbShowed = true
			c.currentResource = "food"
			c.currentResourceIdx = 19
			c.resourceDescriptionIdx = 1

		},
	}
	oxygenIcon := &Button{
		image: c.imgs[2],
		x:     250,
		y:     20,
		onClick: func() {
			c.isDashbShowed = true
			c.currentResource = "oxygen"
			c.currentResourceIdx = 20
			c.resourceDescriptionIdx = 2

		},
	}
	powerIcon := &Button{
		image: c.imgs[3],
		x:     350,
		y:     20,
		onClick: func() {
			c.isDashbShowed = true
			c.currentResource = "power"
			c.currentResourceIdx = 21
			c.resourceDescriptionIdx = 3

		},
	}

	operatorIcon := &Button{
		image: c.imgs[12],
		x:     450,
		y:     20,
		onClick: func() {
			c.operatorStatus = true
			c.logs = append(c.logs, "operator presence confirmed")
		},
	}

	qolIcon := &Button{
		image: c.imgs[22],
		x:     950,
		y:     35,
	}

	c.waterTime = time.Now()
	c.foodTime = time.Now()
	c.oxygenTime = time.Now()
	c.powerTime = time.Now()
	c.operatorTime = time.Now()

	c.icons = append(c.icons, waterIcon, foodIcon, oxygenIcon, powerIcon, qolIcon, operatorIcon)

	// log related
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

	// dashboard related
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

	actionButton1 := &Button{
		image: c.imgs[13],
		x:     (WIDTH/2 + 10) + (c.imgs[16].Bounds().Dx()-c.imgs[13].Bounds().Dx())/2,
		y:     c.imgs[16].Bounds().Dy() + 75,
		text:  "Smaller rations",
		onClick: func() {
			c.resources[c.resourceDescriptionIdx] += 10
			c.KeepInBounds(c.resourceDescriptionIdx)

			c.UpdateLogs(c.currentResource, "up")
			c.qualityOfLiving -= 5
		},
	}

	actionButton2 := &Button{
		image: c.imgs[13],
		x:     (WIDTH/2 + 10) + (c.imgs[16].Bounds().Dx()-c.imgs[13].Bounds().Dx())/2,
		y:     c.imgs[16].Bounds().Dy() + 175,
		text:  "Gather from facility",
		onClick: func() {
			c.resources[c.resourceDescriptionIdx] += 5
			c.KeepInBounds(c.resourceDescriptionIdx)

			c.UpdateLogs(c.currentResource, "up")
			c.qualityOfLiving += 5
		},
	}

	actionButton3 := &Button{
		image: c.imgs[13],
		x:     (WIDTH/2 + 10) + (c.imgs[16].Bounds().Dx()-c.imgs[13].Bounds().Dx())/2,
		y:     c.imgs[16].Bounds().Dy() + 275,
		text:  "Temporary labour",
		onClick: func() {
			c.resources[c.resourceDescriptionIdx] += 5
			c.KeepInBounds(c.resourceDescriptionIdx)
			c.UpdateLogs(c.currentResource, "up")
			c.qualityOfLiving -= 10

		},
	}

	c.buttons = append(c.buttons, logButton, actionButton1, actionButton2, actionButton3)
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

	// quality of life icon
	qolImg, _, _ := ebitenutil.NewImageFromFile(ICON14)

	// buttons
	logImg := ResizeImg(BUTTON4, 260, 52)
	logImg2 := ResizeImg(BUTTON5, 260, 52)

	// log window
	logWindowImg, _, _ := ebitenutil.NewImageFromFile(LOG_WINDOW)

	// dashboards
	dashboardImg1, _, _ := ebitenutil.NewImageFromFile(DASHBOARD1)
	dashboardImg2 := ResizeImg(DASHBOARD2, 405, 750)

	// resources imgs
	resourceImg1 := ResizeImg(IMG1, 300, 300)
	resourceImg2 := ResizeImg(IMG2, 300, 300)
	resourceImg3 := ResizeImg(IMG3, 300, 300)
	resourceImg4 := ResizeImg(IMG4, 300, 300)

	c.imgs = append(c.imgs,
		waterImgG, foodImgG, oxygenImgG, powerImgG,
		waterImgO, foodImgO, oxygenImgO, powerImgO,
		waterImgB, foodImgB, oxygenImgB, powerImgB,
		operatorImg, logImg, logImg2, logWindowImg,
		dashboardImg1, dashboardImg2,
		resourceImg1, resourceImg2, resourceImg3, resourceImg4,
		qolImg)
}

func (c *City) GetResourceDescription() {
	waterD := "Water is crucial for the survival of the Mars colony. It is needed for drinking, growing crops, and maintaining hygiene. Water can be extracted from underground ice deposits or transported from Earth. Efficient recycling systems are also vital to ensure a sustainable water supply."
	foodD := "Food provides the necessary nutrients and energy for the colonists. It can be cultivated in hydroponic farms using Martian soil and water, or delivered from Earth. Developing a self-sufficient food production system is key to the colony's long-term viability."
	oxygenD := "Oxygen is essential for breathing and maintaining life support systems. It can be generated by splitting water molecules through electrolysis or produced by algae in bioreactors. Ensuring a continuous supply of oxygen is critical for the health and safety of all colonists."
	powerD := "Power is needed to run all the colony's systems, including life support, heating, lighting, and machinery. It can be generated from solar panels, nuclear reactors, or other renewable sources. Reliable and efficient power generation is fundamental to the colony's operation and growth."
	c.resourceDescription = append(c.resourceDescription, waterD, foodD, oxygenD, powerD)
}
