package encoder

//go:nosplit
func (e *Encoder) encodeList(data []interface{}) error {
	e.grow(1)
	e.writeByte('l')
	for _, data := range data {
		err := e.encode(data)
		if err != nil {
			return err
		}
	}
	e.grow(1)
	e.writeByte('e')
	return nil
}
