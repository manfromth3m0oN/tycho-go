package buffer

import "github.com/manfromth3m0oN/tycho-go/pkg/errors"

type ReadBuffer struct {
	byteBuf []byte
	index   int
}

func (b ReadBuffer) ReadByte() (byte, error) {
	defer func() { b.index += 1 }()
	if b.byteBuf[b.index] == 0x00 {
		return 0x00, errors.NoMoreBytesError
	}
	return b.byteBuf[b.index], nil
}

func NewReadBuffer(data []byte) ReadBuffer {
	return ReadBuffer{
		byteBuf: data,
		index:   0,
	}
}
