package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Window constants
const (
	screenWidth  = 1280
	screenHeight = 720
	windowTitle  = "Snakers"
	isResizeable = false
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(windowTitle)

	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}

}
