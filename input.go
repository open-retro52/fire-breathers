package firebreathers

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Action int

const (
	Up Action = iota
	Down
	Left
	Right
	Fire
	Nothing
)

type Input interface {
	Action() Action
}

type LocalInput struct {
	up, down, left, right, fire ebiten.Key
}

func NewLocalInput(up, down, left, right, fire ebiten.Key) *LocalInput {
	return &LocalInput{
		up:    up,
		down:  down,
		left:  left,
		right: right,
		fire:  fire,
	}
}

func (l LocalInput) Action() Action {
	return Nothing
}
