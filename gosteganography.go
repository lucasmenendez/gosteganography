package gosteganography

import (
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func parseImage(path string) ([]color.Color, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	pix := []color.Color{}
	img, _, err := image.Decode(reader)

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pix = append(pix, img.At(x, y))
		}
	}
	return pix, err
}

func bytesAvailable(pix []color.Color) int {
	return len(pix) * 3 / 8 // 3 color components divided by number of bits in a byte
}

func num2bin(num uint64) []byte {
	res := make([]byte, 8)
	idx := 0
	for i := num; i != 0; i >>= 1 {
		res[7-idx] = byte(i & 1)
		idx++
	}
	return res
}

func encodeMessage(msg []byte) []byte {
	res := []byte{}
	for _, b := range msg {
		res = append(res, num2bin(uint64(b))...)
	}

	return res
}

// func hideMessage(pix []color.Color, encMessage []byte) []color.Color {}

// func writeImage(path string)

func HideMessage(input, output string, msg []byte) error {
	pix, err := parseImage(input)
	if err != nil {
		return err
	}

	if limit := bytesAvailable(pix); len(msg) > limit {
		return fmt.Errorf("max message size exceeded (%d)", limit)
	}
	encMsg := encodeMessage(msg)
	return nil
}
