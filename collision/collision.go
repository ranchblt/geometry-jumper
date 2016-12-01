/*
Package collision will take 2 image.Image and determine if any of their active pixes are colliding.

This handles offsets for if the pixels are on a Cartesian coordinate system.
*/
package collision

import (
	"image"

	"golang.org/x/image/draw"
)

type Coord interface {
	X() int
	Y() int
}

type coord struct {
	x int
	y int
}

func (c coord) X() int {
	return c.x
}

func (c coord) Y() int {
	return c.y
}

// Hitbox is an image with a center on a Cartesian coordinate system.
type Hitbox struct {
	Image   image.Image
	rgbaImg *image.RGBA
	Center  Coord
}

// IsColliding checks if 2 hitboxs are colliding
func IsColliding(hb1, hb2 *Hitbox) bool {
	return hb1.checkCollision(hb2)
}

// IsCollidingMultiple checks if multiple hitboxes are colliding with the first
func IsCollidingMultiple(hb *Hitbox, boxes []*Hitbox) bool {
	for _, b := range boxes {
		if hb.checkCollision(b) {
			return true
		}
	}
	return false
}

func (hb *Hitbox) width() int {
	return hb.Image.Bounds().Dx()
}

func (hb *Hitbox) height() int {
	return hb.Image.Bounds().Dy()
}

// Check if this shape contains these coords
func (hb *Hitbox) checkIfContains(c Coord) bool {
	minX := hb.Center.X() - hb.width()/2
	maxX := hb.Center.X() + hb.width()/2
	minY := hb.Center.Y() - hb.height()/2
	maxY := hb.Center.Y() + hb.height()/2
	//fmt.Printf("minX: %d, maxX: %d, minY: %d, maxY: %d", minX, maxX, minY, maxY)
	if c.X() > minX && c.X() < maxX && c.Y() > minY && c.Y() < maxY {
		return true
	}
	return false
}

// Based on the center coord convert the pixel in the image to where on the coordinate
// it is.
func (hb *Hitbox) convertPixelToCoord(x, y int) Coord {
	topLeft := coord{
		x: hb.Center.X() - hb.width()/2,
		y: hb.Center.Y() - hb.height()/2,
	}
	topLeft.x = topLeft.x + x
	topLeft.y = topLeft.y + y
	return topLeft
}

// Convert a coordinate location to what number pixel it is.
func (hb *Hitbox) convertCoordToPixel(c Coord) (x, y int) {
	topLeft := coord{
		x: hb.Center.X() - hb.width()/2,
		y: hb.Center.Y() - hb.height()/2,
	}
	px := c.X() - topLeft.x
	py := c.Y() - topLeft.y
	return px, py
}

func (hb *Hitbox) checkCollision(hb2 *Hitbox) bool {
	w := hb.width()
	h := hb.height()

	if hb.rgbaImg == nil {
		hb.rgbaImg = toRGBA(hb.Image)
	}

	if hb2.rgbaImg == nil {
		hb2.rgbaImg = toRGBA(hb2.Image)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			hbPixelColor := hb.rgbaImg.RGBAAt(x, y)

			// Check the alpha, if it is 0 then it is transparent.
			if hbPixelColor.A != 0 {
				c := hb.convertPixelToCoord(x, y)
				if hb2.checkIfContains(c) {
					cx, cy := hb2.convertCoordToPixel(c)

					ccolor := hb2.rgbaImg.RGBAAt(cx, cy)

					// Also check that the second collision Hitbox's alpha
					// is not transparent.
					if ccolor.A != 0 {
						return true
					}
				}
			}
		}
	}

	return false
}

// Convert a image.Image to a RGBA image. This is so we
// can get the color.RGBA
func toRGBA(img image.Image) *image.RGBA {
	switch img.(type) {
	case *image.RGBA:
		return img.(*image.RGBA)
	}
	out := image.NewRGBA(img.Bounds())
	draw.Copy(out, image.Pt(0, 0), img, img.Bounds(), draw.Src, nil)
	return out
}
