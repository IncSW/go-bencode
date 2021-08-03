package decoder

import (
	"bytes"
	"errors"
)

func (d *Decoder) decodeBytes() ([]byte, error) {
	if d.data[d.cursor] < '0' || d.data[d.cursor] > '9' {
		return nil, errors.New("bencode: invalid string field")
	}
	index := bytes.IndexByte(d.data[d.cursor:], ':')
	if index == -1 {
		return nil, errors.New("bencode: invalid string field")
	}
	index += d.cursor
	stringLength, err := d.parseInt(d.data[d.cursor:index])
	if err != nil {
		return nil, err
	}
	index += 1
	endIndex := index + int(stringLength)
	if endIndex > d.length {
		return nil, errors.New("bencode: not a valid bencoded string")
	}
	value := d.data[index:endIndex]
	d.cursor = endIndex
	return value, nil
}
