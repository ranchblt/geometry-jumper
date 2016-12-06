package menu

import (
	"geometry-jumper/keyboard"
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
)

type Menu interface {
	Selected() string
	Draw(*ebiten.Image)
	Update()
}

type Option struct {
	Text string
}

const size = 24
const dpi = 72

type Regular struct {
	BackgroundImage *ebiten.Image
	Options         []*Option
	KeyboardWrapper *keyboard.KeyboardWrapper
	selected        int
	Height          int
	Width           int
	Font            *truetype.Font
}

// Selected gives the text of the currently selected option
func (r *Regular) Selected() string {
	return r.Options[r.selected].Text
}

// SelectedIncrease moves which option the user has selected up one
func (r *Regular) selectedIncrease() {
	if r.selected+1 > len(r.Options)-1 {
		r.selected = 0
		return
	}
	r.selected++
}

// SelectedDecrease moves which option the user has selected down one
func (r *Regular) selectedDecrease() {
	if r.selected-1 < 0 {
		r.selected = len(r.Options) - 1
		return
	}
	r.selected--
}

func (r *Regular) Update() {
	if r.KeyboardWrapper.KeyPushed(ebiten.KeyUp) {
		r.selectedIncrease()
	}

	if r.KeyboardWrapper.KeyPushed(ebiten.KeyDown) {
		r.selectedDecrease()
	}
}

func (r *Regular) Draw(screen *ebiten.Image) {
	screen.DrawImage(r.BackgroundImage, &ebiten.DrawImageOptions{
		ImageParts: r,
	})

	textImage, _ := ebiten.NewImage(r.Width, r.Height, ebiten.FilterNearest)
	selectedImage, _ := ebiten.NewImage(r.Width, r.Height, ebiten.FilterNearest)
	w, h := textImage.Size()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	dst2 := image.NewRGBA(image.Rect(0, 0, w, h))

	d := &font.Drawer{
		Dst: dst,
		Src: image.White,
		Face: truetype.NewFace(r.Font, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}

	d2 := &font.Drawer{
		Dst: dst2,
		Src: image.White,
		Face: truetype.NewFace(r.Font, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}

	y := r.Height / 2
	for i, o := range r.Options {
		drawer := d
		if i == r.selected {
			drawer = d2
		}

		s := font.MeasureString(d.Face, o.Text)
		drawer.Dot = fixed.P(r.Width/2-s.Round()/2, y)
		drawer.DrawString(o.Text)
		y += size
	}

	textImage.ReplacePixels(dst.Pix)
	selectedImage.ReplacePixels(dst2.Pix)

	screen.DrawImage(textImage, &ebiten.DrawImageOptions{})

	cm := ebiten.ColorM{}
	cm.Scale(255, 0, 0, 100)

	screen.DrawImage(selectedImage, &ebiten.DrawImageOptions{
		ColorM: cm,
	})
}

func (r *Regular) Len() int {
	return 1
}

func (r *Regular) Src(i int) (x0, y0, x1, y1 int) {
	x, y := r.BackgroundImage.Size()
	return 0, 0, x, y
}

func (r *Regular) Dst(i int) (x0, y0, x1, y1 int) {
	return 0, 0, r.Width, r.Height
}
