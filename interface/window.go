// _interface/window.go
package interfacepkg

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Window struct {
	background *ebiten.Image
}

func (w *Window) Init() {
	w.background, _, _ = ebitenutil.NewImageFromFile(BG)
}

func (w *Window) Update() error {
}

func (w *Window) Draw(screen *ebiten.Image) {
	opBG := AdjustSize(w.background, 2, 2)

}
