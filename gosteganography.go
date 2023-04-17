package gosteganography

import "fmt"

func HideMessage(input, output string, msg []byte) (int, error) {
	img, err := readFile(input)
	if err != nil {
		return 0, err
	}

	if limit := img.bytesAvailable(); len(msg) > limit {
		return 0, fmt.Errorf("max message size exceeded (%d)", limit)
	}
	encMsg := encodeMessage(msg)
	img.hide(encMsg)

	return len(encMsg), img.writeFile(output)
}
