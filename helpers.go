package gosteganography

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
