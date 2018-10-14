package bencode

import (
	"errors"
	"strconv"
)

func readUntil(data []byte, symbol byte) ([]byte, int, bool) {
	for i, b := range data {
		if b == symbol {
			return data[:i], i, true
		}
	}

	return nil, 0, false
}

func Unmarshal(data []byte) (interface{}, error) {
	result, _, err := unmarshal(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func unmarshal(data []byte) (interface{}, int, error) {
	switch data[0] {
	case 'i':
		integerBuffer, length, ok := readUntil(data[1:], 'e')
		if !ok {
			return nil, 0, errors.New("bencode: invalid integer field")
		}

		integer, err := strconv.ParseInt(string(integerBuffer), 10, 64)
		if err != nil {
			return nil, 0, err
		}

		return integer, length + 2, nil

	case 'l':
		list := []interface{}{}
		totalLength := 2
		data = data[1:]
		for {
			if len(data) == 0 {
				return nil, 0, errors.New("bencode: invalid list field")
			}

			if data[0] == 'e' {
				return list, totalLength, nil
			}

			value, length, err := unmarshal(data)
			if err != nil {
				return nil, 0, err
			}

			list = append(list, value)
			data = data[length:]
			totalLength += length
		}

	case 'd':
		dictionary := map[string]interface{}{}
		totalLength := 2
		data = data[1:]
		for {
			if len(data) == 0 {
				return nil, 0, errors.New("bencode: invalid dictionary field")
			}

			if data[0] == 'e' {
				return dictionary, totalLength, nil
			}

			value, length, err := unmarshal(data)
			if err != nil {
				return nil, 0, err
			}

			key, ok := value.([]byte)
			if !ok {
				return nil, 0, errors.New("bencode: non-string dictionary key")
			}

			data = data[length:]
			totalLength += length
			value, length, err = unmarshal(data)
			if err != nil {
				return nil, 0, err
			}

			dictionary[string(key)] = value
			data = data[length:]
			totalLength += length
		}

	default:
		stringLengthBuffer, length, ok := readUntil(data, ':')
		if !ok {
			return nil, 0, errors.New("bencode: invalid string field")
		}

		stringLength, err := strconv.ParseInt(string(stringLengthBuffer), 10, 64)
		if err != nil {
			return nil, 0, err
		}

		endPosition := length + 1 + int(stringLength)
		if endPosition > len(data) {
			return nil, 0, errors.New("bencode: not a valid bencoded string")
		}

		return data[length+1 : endPosition], endPosition, nil
	}
}
