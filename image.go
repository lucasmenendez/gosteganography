package gosteganography

import (
	"image"
	"image/color"
	"image/png"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type pixel struct {
	x, y  int
	color color.Color
}

type img struct {
	original image.Image
	bounds   image.Rectangle
	imgType  string
	pixels   []*pixel
}

func readFile(path string) (*img, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	original, itype, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	res := &img{
		original: original,
		bounds:   original.Bounds(),
		imgType:  itype,
		pixels:   []*pixel{},
	}

	for y := res.bounds.Min.Y; y < res.bounds.Max.Y; y++ {
		for x := res.bounds.Min.X; x < res.bounds.Max.X; x++ {
			res.pixels = append(res.pixels, &pixel{x, y, original.At(x, y)})
		}
	}

	return res, err
}

func (i *img) bytesAvailable() int {
	return len(i.pixels) * 3 / 8 // 3 color components divided by number of bits in a byte
}

func (i *img) hide(bmsg []byte) {
	newPixels := make([]*pixel, len(i.pixels))

	idx := 0
	mgsLen := len(bmsg)
	for i, p := range i.pixels {
		r, g, b, a := p.color.RGBA()

		if idx < mgsLen {
			r += uint32(bmsg[idx])
		}
		if idx+1 < mgsLen {
			g += uint32(bmsg[idx+1])
		}
		if idx+2 < mgsLen {
			b += uint32(bmsg[idx+2])
		}

		newPixels[i] = &pixel{
			x: p.x, y: p.y,
			color: color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)},
		}
		idx += 3
	}
	i.pixels = newPixels
}

func (i *img) writeFile(path string) error {
	var newImage = image.NewRGBA(i.bounds)
	for _, pix := range i.pixels {
		newImage.Set(pix.x, pix.y, pix.color)
	}

	output, err := os.Create(path)
	if err != nil {
		return err
	}

	switch i.imgType {
	case "png":
		return png.Encode(output, newImage)
	}

	return nil
}
