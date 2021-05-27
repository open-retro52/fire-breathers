package standalone

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	firebreathers "github.com/open-retro52/fire-breathers"
)

func Run() {
	player1 := firebreathers.NewLocalInput(ebiten.KeyW, ebiten.KeyS, ebiten.KeyA, ebiten.KeyD, ebiten.KeySpace)
	player2 := firebreathers.NewLocalInput(ebiten.KeyNumpad8, ebiten.KeyNumpad5, ebiten.KeyNumpad4, ebiten.KeyNumpad6, ebiten.KeyNumpad0)

	game, err := firebreathers.NewGame(960, 540, player1, player2)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Fire-Breathers")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
