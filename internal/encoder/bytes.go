package encoder

//go:nosplit
func (e *Encoder) encodeBytes(data []byte) {
	dataLength := len(data)
	e.grow(dataLength + 23)
	e.writeInt(int64(len(data)))
	e.writeByte(':')
	e.write(data)
}
