// Copyright (c) 2023, Lucas Menendez <hi@lucasmenendez.me>
// See LICENSE for licensing information

package gosteganography

const bitsPerByte = 8

func num2bin(num uint) []uint {
	res := []uint{}
	for i := num; i != 0; i >>= 1 {
		res = append(res, i&1)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func bin2num(bin []uint) uint {
	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	result := uint(0)
	for i := len(bin); i > 0; i-- {
		result = (result << 1) | bin[i-1]
	}
	return result
}

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

func encodeMessage(msg []byte) []uint {
	res := []uint{}
	for _, ibyte := range msg {
		ibin := resize(num2bin(uint(ibyte)), bitsPerByte)
		res = append(res, ibin...)
	}

	return res
}

func decodeMessage(bin []uint) []byte {
	res := []byte{}
	for _, ibinbyte := range split(bin, bitsPerByte) {
		res = append(res, byte(bin2num(ibinbyte)))
	}

	return res
}
