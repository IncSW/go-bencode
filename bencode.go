package bencode

import (
	"github.com/IncSW/go-bencode/internal/decoder"
	"github.com/IncSW/go-bencode/internal/encoder"
)

func MarshalTo(dst []byte, data interface{}) ([]byte, error) {
	var e encoder.Encoder
	return e.EncodeTo(dst, data)
}

func Marshal(data interface{}) ([]byte, error) {
	var e encoder.Encoder
	return e.EncodeTo(nil, data)
}

func Unmarshal(data []byte) (interface{}, error) {
	return decoder.Unmarshal(data)
}
