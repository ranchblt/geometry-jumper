package game

import (
	"bytes"
	"image"
	"time"

	"github.com/uber-go/zap"

	"github.com/ranchblt/geometry-jumper/resource"

	"golang.org/x/image/draw"
)

func openImage(path string) (image.Image, error) {
	b, err := resource.Asset(path)
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	return image, nil
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	logger.Info("Time Track",
		zap.String("Name", name),
		zap.Duration("Elapsed", elapsed),
	)
}

func toRGBA(img image.Image) *image.RGBA {
	switch img.(type) {
	case *image.RGBA:
		return img.(*image.RGBA)
	}
	out := image.NewRGBA(img.Bounds())
	draw.Copy(out, image.Pt(0, 0), img, img.Bounds(), draw.Src, nil)
	return out
}
