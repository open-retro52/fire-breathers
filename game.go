package firebreathers

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player        [2]Input
	width, height int
}

// NewGame generates a new Game object.
func NewGame(width, height int, p1, p2 Input) (*Game, error) {
	g := &Game{
		player: [2]Input{
			p1,
			p2,
		},
		width:  width,
		height: height,
	}

	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

// Update updates the current game state.
func (g *Game) Update() error {
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
}
