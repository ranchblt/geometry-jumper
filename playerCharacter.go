package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type PlayerCharacter struct {
	name  string
	Image *ebiten.Image
}

func (pc *PlayerCharacter) Update() error {
	if keyboardWrapper.KeyPushed(ebiten.KeySpace) {
		fmt.Print("you pushed space")
	}
	return nil
}

func (p *PlayerCharacter) Len() int {
	w, h := p.Image.Size()
	return (screenWidth/w + 1) * (screenHeight/h + 2)
}

func (p *PlayerCharacter) Dst(i int) (x0, y0, x1, y1 int) {
	return 20, 20, 60, 60
}

func (p *PlayerCharacter) Src(i int) (x0, y0, x1, y1 int) {
	w, h := p.Image.Size()
	return 0, 0, w, h
}
