package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var currentState int

type Window struct {
	background *ebiten.Image
	buttons    []*Button
	gameLogic  []Game
}

func (w *Window) Init() {
	w.background, _, _ = ebitenutil.NewImageFromFile(BG)
	buttonImg, _, _ := ebitenutil.NewImageFromFile(BUTTON1)
	buttonImgLeft, _, _ := ebitenutil.NewImageFromFile(BUTTON2)
	buttonImgRight, _, _ := ebitenutil.NewImageFromFile(BUTTON3)
	cityImg, _, _ := ebitenutil.NewImageFromFile(CITY)
	gardensImg, _, _ := ebitenutil.NewImageFromFile(GARDENS)
	reactorImg, _, _ := ebitenutil.NewImageFromFile(REACTOR)
	transportationImg, _, _ := ebitenutil.NewImageFromFile(TRANSPORTATION)

	startButton := &Button{
		image: buttonImg,
		x:     (WIDTH - buttonImg.Bounds().Dx()) / 2,
		y:     HEIGHT/2 + 120,
		text:  "Start",
		onClick: func() {
			currentState = StateCity
		},
	}

	exitButton := &Button{
		image: buttonImg,
		x:     (WIDTH - buttonImg.Bounds().Dx()) / 2,
		y:     HEIGHT/2 + 230,
		text:  "Back to Earth",
		onClick: func() {
			currentState = StateExit
		},
	}

	leftButton := &Button{
		image:   buttonImgLeft,
		x:       50,
		y:       HEIGHT / 2,
		text:    "",
		onClick: func() { currentState = StateGardens },
	}
	rightButton := &Button{
		image:   buttonImgRight,
		x:       WIDTH - buttonImgRight.Bounds().Dx() - 50,
		y:       HEIGHT / 2,
		text:    "",
		onClick: func() { currentState = StateTransportation },
	}

	w.buttons = append(w.buttons, startButton, exitButton, leftButton, rightButton)

	city := &City{
		bg:              cityImg,
		water:           100,
		food:            100,
		oxygen:          100,
		power:           100,
		qualityOfLiving: 100,
	}
	city.Init()

	gardens := &Gardens{
		bg: gardensImg,
	}
	reactor := &Reactor{
		bg: reactorImg,
	}
	transportation := &Transportation{
		bg: transportationImg,
	}

	w.gameLogic = append(w.gameLogic, city, gardens, reactor, transportation)
}

func (w *Window) Update() error {
	switch currentState {
	case StateMenu:
		w.buttons[0].Update()
		w.buttons[1].Update()

	case StateCity:
		w.buttons[2].onClick = func() { currentState = StateGardens }
		w.buttons[3].onClick = func() { currentState = StateTransportation }
		w.buttons[2].Update()
		w.buttons[3].Update()
		if city, ok := w.gameLogic[0].(*City); ok {
			city.UpdateResources()
		}

	case StateGardens:
		w.buttons[3].onClick = func() { currentState = StateCity }
		w.buttons[3].Update()

	case StateReactor:
		w.buttons[2].onClick = func() { currentState = StateTransportation }
		w.buttons[2].Update()

	case StateTransportation:
		w.buttons[2].onClick = func() { currentState = StateCity }
		w.buttons[3].onClick = func() { currentState = StateReactor }
		w.buttons[2].Update()
		w.buttons[3].Update()

	case StateExit:
		return ebiten.Termination
	}

	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	switch currentState {
	case StateMenu:
		opBG := AdjustSize(w.background, 2, 2)
		screen.DrawImage(w.background, opBG)
		w.buttons[0].Draw(screen)
		w.buttons[1].Draw(screen)

	case StateCity:
		if city, ok := w.gameLogic[0].(*City); ok {
			city.Draw(screen)
		}
		w.buttons[2].Draw(screen)
		w.buttons[3].Draw(screen)

	case StateGardens:
		if gardens, ok := w.gameLogic[1].(*Gardens); ok {
			gardens.Draw(screen)
		}
		w.buttons[3].Draw(screen)

	case StateReactor:
		if reactor, ok := w.gameLogic[2].(*Reactor); ok {
			reactor.Draw(screen)
		}
		w.buttons[2].Draw(screen)

	case StateTransportation:
		if transportation, ok := w.gameLogic[3].(*Transportation); ok {
			transportation.Draw(screen)
		}
		w.buttons[2].Draw(screen)
		w.buttons[3].Draw(screen)

	}
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}
