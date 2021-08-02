package bencode

import (
	"github.com/IncSW/go-bencode/internal/decoder"
	"github.com/IncSW/go-bencode/internal/encoder"
)

func MarshalTo(dst []byte, data interface{}) ([]byte, error) {
	return encoder.MarshalTo(dst, data)
}

func Marshal(data interface{}) ([]byte, error) {
	return encoder.MarshalTo(make([]byte, 512), data)
}

func Unmarshal(data []byte) (interface{}, error) {
	return decoder.Unmarshal(data)
}
