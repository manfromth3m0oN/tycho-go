package buffer

type WriteBuffer struct {
	byteBuf []byte
}

func (b WriteBuffer) WriteByte(p byte) {
	b.byteBuf = append(b.byteBuf, p)
}
