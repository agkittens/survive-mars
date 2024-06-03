package main

// window consts
const WIDTH, HEIGHT = 1024, 750
const TPS = 60

// imgs
const BG = "interface/assets/bg.jpg"
const CITY = "interface/assets/city.jpg"
const GARDENS = "interface/assets/gardens.jpg"
const REACTOR = "interface/assets/reactor.jpg"
const TRANSPORTATION = "interface/assets/transportation.jpg"

// buttons imgs
const BUTTON1 = "interface/assets/Icons/Arrows.png"
const BUTTON2 = "interface/assets/Icons/ArrowsLeft2.png"
const BUTTON3 = "interface/assets/Icons/ArrowsRight2.png"

// window
const (
	StateMenu = iota
	StateCity
	StateGardens
	StateReactor
	StateTransportation
	StateExit
)
