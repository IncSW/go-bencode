package encoder

import (
	"fmt"
	"unsafe"

	"github.com/IncSW/go-bencode/internal"
)

//go:linkname memmove runtime.memmove
//go:noescape
func memmove(to unsafe.Pointer, from unsafe.Pointer, n uintptr)

type sliceHeader struct {
	data unsafe.Pointer
	len  int
	cap  int
}

type Encoder struct {
	buffer []byte
	length int
	offset int
}

//go:nosplit
func (e *Encoder) grow(neededLength int) {
	availableLength := e.length - e.offset
	if availableLength >= neededLength {
		return
	}
	if e.length == 0 {
		if neededLength < 16 {
			neededLength = 16
		}
		e.length = neededLength
		availableLength = neededLength
	} else {
		for availableLength < neededLength {
			e.length += e.length
			availableLength = e.length - e.offset
		}
	}
	buffer := make([]byte, e.length)
	memmove(
		unsafe.Pointer(uintptr((*sliceHeader)(unsafe.Pointer(&buffer)).data)),
		(*sliceHeader)(unsafe.Pointer(&e.buffer)).data,
		uintptr(e.offset),
	)
	e.buffer = buffer
}

//go:nosplit
func (e *Encoder) write(data []byte) {
	length := len(data)
	memmove(
		unsafe.Pointer(uintptr((*sliceHeader)(unsafe.Pointer(&e.buffer)).data)+uintptr(e.offset)),
		(*sliceHeader)(unsafe.Pointer(&data)).data,
		uintptr(length),
	)
	e.offset += length
}

//go:nosplit
func (e *Encoder) writeByte(data byte) {
	*(*byte)(unsafe.Pointer(uintptr((*sliceHeader)(unsafe.Pointer(&e.buffer)).data) + uintptr(e.offset))) = data
	e.offset++
}

//go:nosplit
func (e *Encoder) EncodeTo(dst []byte, data interface{}) ([]byte, error) {
	if cap(dst) > len(dst) {
		dst = dst[:cap(dst)]
	} else if len(dst) == 0 {
		dst = make([]byte, 512)
	}
	e.buffer = dst
	e.length = cap(dst)
	err := e.encode(data)
	if err != nil {
		return nil, err
	}
	return e.buffer[:e.offset], nil
}

//go:nosplit
func (e *Encoder) encode(data interface{}) error {
	switch value := data.(type) {
	case int64:
		e.encodeInt(value)
	case int32:
		e.encodeInt(int64(value))
	case int16:
		e.encodeInt(int64(value))
	case int8:
		e.encodeInt(int64(value))
	case int:
		e.encodeInt(int64(value))
	case uint64:
		e.encodeInt(int64(value))
	case uint32:
		e.encodeInt(int64(value))
	case uint16:
		e.encodeInt(int64(value))
	case uint8:
		e.encodeInt(int64(value))
	case uint:
		e.encodeInt(int64(value))
	case []byte:
		e.encodeBytes(value)
	case string:
		e.encodeBytes(internal.S2B(value))
	case []interface{}:
		return e.encodeList(value)
	case map[string]interface{}:
		return e.encodeDictionary(value)
	default:
		return fmt.Errorf("bencode: unsupported type: %T", value)
	}
	return nil
}
