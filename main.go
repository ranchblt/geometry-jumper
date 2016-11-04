package main

import (
        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
        ebitenutil.DebugPrint(screen, "Hello world!")
        return nil
}

func main() {
        ebiten.Run(update, 320, 240, 2, "Hello world!")
}