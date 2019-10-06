package bencode

import (
	"bytes"
	"errors"
	"strconv"
)

func Unmarshal(data []byte) (interface{}, error) {
	return (&unmarshaler{
		data:   data,
		length: len(data),
	}).unmarshal()
}

type unmarshaler struct {
	data   []byte
	length int
	cursor int
}

func (u *unmarshaler) unmarshal() (interface{}, error) {
	switch u.data[u.cursor] {
	case 'i':
		u.cursor += 1
		index := bytes.IndexByte(u.data[u.cursor:], 'e')
		if index == -1 {
			return nil, errors.New("bencode: invalid integer field")
		}
		index += u.cursor
		integer, err := strconv.ParseInt(b2s(u.data[u.cursor:index]), 10, 64)
		if err != nil {
			return nil, err
		}
		u.cursor = index + 1
		return integer, nil

	case 'l':
		u.cursor += 1
		list := []interface{}{}
		for {
			if u.cursor == u.length {
				return nil, errors.New("bencode: invalid list field")
			}
			if u.data[u.cursor] == 'e' {
				u.cursor += 1
				return list, nil
			}
			value, err := u.unmarshal()
			if err != nil {
				return nil, err
			}
			list = append(list, value)
		}

	case 'd':
		u.cursor += 1
		dictionary := map[string]interface{}{}
		for {
			if u.cursor == u.length {
				return nil, errors.New("bencode: invalid dictionary field")
			}
			if u.data[u.cursor] == 'e' {
				u.cursor += 1
				return dictionary, nil
			}
			value, err := u.unmarshal()
			if err != nil {
				return nil, err
			}
			key, ok := value.([]byte)
			if !ok {
				return nil, errors.New("bencode: non-string dictionary key")
			}
			value, err = u.unmarshal()
			if err != nil {
				return nil, err
			}
			dictionary[b2s(key)] = value
		}

	default:
		index := bytes.IndexByte(u.data[u.cursor:], ':')
		if index == -1 {
			return nil, errors.New("bencode: invalid string field")
		}
		index += u.cursor
		stringLength, err := strconv.ParseInt(b2s(u.data[u.cursor:index]), 10, 64)
		if err != nil {
			return nil, err
		}
		index += 1
		endIndex := index + int(stringLength)
		if endIndex > u.length {
			return nil, errors.New("bencode: not a valid bencoded string")
		}
		value := u.data[index:endIndex]
		u.cursor = endIndex
		return value, nil
	}
}
