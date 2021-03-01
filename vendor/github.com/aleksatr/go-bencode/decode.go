package bencode

import (
	"bytes"
	"errors"
	"strconv"
)

var ErrBadEncoding = errors.New("bad encoding")
var ErrInvalidInput = errors.New("invalid input")

var StringDelimiter = []byte(":")

const (
	DigitLow        = byte('0')
	DigitHigh       = byte('9')
	IntegerStart    = byte('i')
	IntegerEnd      = byte('e')
	ListStart       = byte('l')
	ListEnd         = byte('e')
	DictionaryStart = byte('d')
	DictionaryEnd   = byte('e')
)

// Decode will return either string, int64, []interface{} or map[string]interface{},
// interface{} being anything of these 4 types
func Decode(data []byte) (interface{}, error) {
	obj, _, err := decodeNextObject(data)
	return obj, err
}

// decodeNextObject returns value of the next object, offset of the value after it and error
func decodeNextObject(data []byte) (interface{}, int, error) {
	if len(data) == 0 {
		return nil, 0, ErrInvalidInput
	}

	c := data[0]
	switch {
	case c >= DigitLow && c <= DigitHigh:
		return decodeString(data)
	case c == IntegerStart:
		return decodeInteger(data)
	case c == ListStart:
		return decodeList(data)
	case c == DictionaryStart:
		return decodeDictionary(data)
	}

	return nil, 0, ErrBadEncoding
}

func decodeString(data []byte) (string, int, error) {
	c := data[0]
	if c < DigitLow || c > DigitHigh {
		return "", 0, ErrBadEncoding
	}

	parts := bytes.SplitN(data, StringDelimiter, 2)
	switch {
	case len(parts) != 2 && c != DigitLow:
		return "", 0, ErrBadEncoding
	case c == DigitLow:
		return "", 2, nil
	}

	// TODO: maybe strconv.ParseInt(string(parts[0]), 10, 64) if we want int64
	l, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return "", 0, err
	}

	if len(parts[1]) < l {
		return "", 0, ErrBadEncoding
	}

	return string(parts[1][:l]), len(parts[0]) + 1 + l, nil
}

func decodeInteger(data []byte) (int64, int, error) {
	c := data[0]
	if c != IntegerStart {
		return 0, 0, ErrBadEncoding
	}

	e := bytes.IndexByte(data, IntegerEnd)
	if e < 2 {
		return 0, 0, ErrBadEncoding
	}

	n, err := strconv.ParseInt(string(data[1:e]), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return n, e + 1, err
}

func decodeList(data []byte) ([]interface{}, int, error) {
	c := data[0]
	if c != ListStart {
		return nil, 0, ErrBadEncoding
	}

	var list []interface{}
	p := 1
	for {
		if data[p] == ListEnd {
			break
		}

		obj, d, err := decodeNextObject(data[p:])
		if err != nil {
			return nil, 0, err
		}

		list = append(list, obj)
		p += d
	}

	return list, p + 1, nil
}

func decodeDictionary(data []byte) (map[string]interface{}, int, error) {
	c := data[0]
	if c != DictionaryStart {
		return nil, 0, ErrBadEncoding
	}

	dict := make(map[string]interface{})
	lastKey := ""
	p := 1 // offset in the input slice
	for {
		if data[p] == DictionaryEnd {
			break
		}

		// key must be a string
		key, i, err := decodeString(data[p:])
		if err != nil {
			return nil, 0, err
		}

		// keys must be ordered
		if p > 1 && lastKey >= key {
			return nil, 0, ErrBadEncoding
		}

		lastKey = key
		p += i

		// value
		val, j, err := decodeNextObject(data[p:])
		if err != nil {
			return nil, 0, err
		}

		p += j

		dict[key] = val
	}

	return dict, p + 1, nil
}
