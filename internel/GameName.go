package internel

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "More to come...")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func Loop() {
	game := &Game{}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("More to come...")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}