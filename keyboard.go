package main

import "github.com/hajimehoshi/ebiten"

type KeyboardWrapper struct {
	keyState map[ebiten.Key]int
	keyNames map[ebiten.Key]string
}

func NewKeyboardWrapper() *KeyboardWrapper {
	var kw = &KeyboardWrapper{
		keyState: map[ebiten.Key]int{},
		keyNames: map[ebiten.Key]string{
			ebiten.KeyBackspace: "BS",
			ebiten.KeyComma:     ",",
			ebiten.KeyDelete:    "Del",
			ebiten.KeyEnter:     "Enter",
			ebiten.KeyEscape:    "Esc",
			ebiten.KeyPeriod:    ".",
			ebiten.KeySpace:     "Space",
			ebiten.KeyTab:       "Tab",

			// Arrows
			ebiten.KeyDown:  "Down",
			ebiten.KeyLeft:  "Left",
			ebiten.KeyRight: "Right",
			ebiten.KeyUp:    "Up",

			// Mods
			ebiten.KeyShift:   "Shift",
			ebiten.KeyControl: "Ctrl",
			ebiten.KeyAlt:     "Alt",
		},
	}
	return kw
}

func (kw *KeyboardWrapper) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func (kw *KeyboardWrapper) KeyPushed(key ebiten.Key) bool {
	return kw.keyState[ebiten.KeySpace] == 1
}

func (kw *KeyboardWrapper) Update() {
	// A through Z
	for c := 'A'; c <= 'Z'; c++ {
		var key = ebiten.Key(c) - 'A' + ebiten.KeyA
		if kw.IsKeyPressed(key) {
			kw.keyState[key]++
		} else {
			kw.keyState[key] = 0
		}
	}

	// numerics
	for i := 1; i <= 12; i++ {
		var key = ebiten.Key(i) + ebiten.KeyF1 - 1
		if kw.IsKeyPressed(key) {
			kw.keyState[key]++
		} else {
			kw.keyState[key] = 0
		}
	}

	// and the weird ones
	for key := range kw.keyNames {
		if kw.IsKeyPressed(key) {
			kw.keyState[key]++
		} else {
			kw.keyState[key] = 0
		}
	}
}
