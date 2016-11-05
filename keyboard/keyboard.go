package keyboard

import "github.com/hajimehoshi/ebiten"

type KeyboardWrapper struct {
	keyState map[ebiten.Key]int
	keys     []ebiten.Key
}

func NewKeyboardWrapper() *KeyboardWrapper {
	var kw = &KeyboardWrapper{
		keyState: map[ebiten.Key]int{},
		keys: []ebiten.Key{
			ebiten.Key0,
			ebiten.Key1,
			ebiten.Key2,
			ebiten.Key3,
			ebiten.Key4,
			ebiten.Key5,
			ebiten.Key6,
			ebiten.Key7,
			ebiten.Key8,
			ebiten.Key9,
			ebiten.KeyA,
			ebiten.KeyB,
			ebiten.KeyC,
			ebiten.KeyD,
			ebiten.KeyE,
			ebiten.KeyF,
			ebiten.KeyG,
			ebiten.KeyH,
			ebiten.KeyI,
			ebiten.KeyJ,
			ebiten.KeyK,
			ebiten.KeyL,
			ebiten.KeyM,
			ebiten.KeyN,
			ebiten.KeyO,
			ebiten.KeyP,
			ebiten.KeyQ,
			ebiten.KeyR,
			ebiten.KeyS,
			ebiten.KeyT,
			ebiten.KeyU,
			ebiten.KeyV,
			ebiten.KeyW,
			ebiten.KeyX,
			ebiten.KeyY,
			ebiten.KeyZ,
			ebiten.KeyAlt,
			ebiten.KeyBackspace,
			ebiten.KeyCapsLock,
			ebiten.KeyComma,
			ebiten.KeyControl,
			ebiten.KeyDelete,
			ebiten.KeyDown,
			ebiten.KeyEnd,
			ebiten.KeyEnter,
			ebiten.KeyEscape,
			ebiten.KeyF1,
			ebiten.KeyF2,
			ebiten.KeyF3,
			ebiten.KeyF4,
			ebiten.KeyF5,
			ebiten.KeyF6,
			ebiten.KeyF7,
			ebiten.KeyF8,
			ebiten.KeyF9,
			ebiten.KeyF10,
			ebiten.KeyF11,
			ebiten.KeyF12,
			ebiten.KeyHome,
			ebiten.KeyInsert,
			ebiten.KeyLeft,
			ebiten.KeyPageDown,
			ebiten.KeyPageUp,
			ebiten.KeyPeriod,
			ebiten.KeyRight,
			ebiten.KeyShift,
			ebiten.KeySpace,
			ebiten.KeyTab,
			ebiten.KeyUp,
			ebiten.KeyMax,
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
	for _, key := range kw.keys {
		if kw.IsKeyPressed(key) {
			kw.keyState[key]++
		} else {
			kw.keyState[key] = 0
		}
	}
}
