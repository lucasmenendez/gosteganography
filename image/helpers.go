// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package image

// bitsPerByte constant contains the number of bits in a byte
const bitsPerByte = 8

// num2bin function gets the binary representation of the number provided and
// returns it in a slice of unsigned integers.
func num2bin(num uint) []uint {
	res := []uint{}
	// iterate over the number of bits using the bitwise right shift operator
	// to move over the number in binary
	for i := num; i != 0; i >>= 1 {
		// append the last bit using the bitwise and operator
		res = append(res, i&1)
	}
	// reverse the binary representation
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// bin2num function gets the decimal representation of the slice of bits
// (uint's) provided and returns it in a unsigned integer.
func bin2num(bin []uint) uint {
	// reverse back the binary representation
	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}
	result := uint(0)
	// rebuild the number using the bitwise left shift and or operators
	for i := len(bin); i > 0; i-- {
		result = (result << 1) | bin[i-1]
	}
	return result
}

// resize function returns the given bit slice modified to fit the given size:
//   - If the given size is less than the length of the input, it truncates the
//     input.
//   - If the size is greater than the length of the input, padding is added.
//   - If the given size is the current length, the input is returned
//     unmodified.
func resize(bits []uint, to int) []uint {
	n := to - len(bits)
	if n > 0 {
		padding := make([]uint, n)
		return append(padding, bits...)
	} else if n < 0 {
		return bits[:to]
	}
	return bits
}

// split function returns the given slice of unsigned ints, split into groups
// of the given size.
func split(input []uint, size int) [][]uint {
	var chunks [][]uint
	for {
		if len(input) == 0 {
			break
		}
		if len(input) < size {
			size = len(input)
		}
		chunks = append(chunks, input[0:size])
		input = input[size:]
	}
	return chunks
}

// encodeMessage function iterates over the bytes of the provided slice, gets
// its binary representation and returns them concatenated.
func encodeMessage(msg []byte) []uint {
	res := []uint{}
	for _, ibyte := range msg {
		ibin := resize(num2bin(uint(ibyte)), bitsPerByte)
		res = append(res, ibin...)
	}
	return res
}

// decodeMessage function splits the given binary slice into bits of bytes,
// gets the decimal byte code of them and returns them in a slice of bytes.
func decodeMessage(bin []uint) []byte {
	res := []byte{}
	for _, ibinbyte := range split(bin, bitsPerByte) {
		res = append(res, byte(bin2num(ibinbyte)))
	}
	return res
}
