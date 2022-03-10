package length

import (
	"github.com/manfromth3m0oN/tycho-go/pkg/errors"
)

func ReadLength(buffer []byte) (int, error) {
	var num uint64
	for i, byte := range buffer {
		num |= (uint64((byte & 0x7F)) << (7 * i))
		if byte&0x80 == 0 {
			return int(num), nil
		}
	}
	return 0, errors.InvalidLengthError
}

func ()

func WriteLength(length int) ([]byte, error) {
	buf := make([]byte, 0)
	for {
		byte := byte(length & 0x7F)
		cont := length >> 7
		if cont == 0 {
			buf = append(buf, byte)
			return buf, nil
		} else {
			buf = append(buf, byte)
		}
	}
}
