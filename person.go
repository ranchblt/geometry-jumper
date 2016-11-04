package main

import (
	"github.com/hajimehoshi/ebiten"
)

var personImage *ebiten.Image

type personImageParts struct {
	image *ebiten.Image
}

func (p *personImageParts) Len() int {
	w, h := p.image.Size()
	return (screenWidth/w + 1) * (screenHeight/h + 2)
}

func (p *personImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	return 20, 20, 60, 60
}

func (p *personImageParts) Src(i int) (x0, y0, x1, y1 int) {
	w, h := p.image.Size()
	return 0, 0, w, h
}
