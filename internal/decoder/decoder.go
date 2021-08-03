package decoder

import (
	"errors"

	"github.com/IncSW/go-bencode/internal"
)

type Decoder struct {
	data   []byte
	length int
	cursor int
}

func (d *Decoder) Decode(data []byte) (interface{}, error) {
	d.data = data
	d.length = len(data)
	return d.decode()
}

func (d *Decoder) decode() (interface{}, error) {
	switch d.data[d.cursor] {
	case 'i':
		return d.decodeInt()
	case 'l':
		d.cursor += 1
		list := []interface{}{}
		for {
			if d.cursor == d.length {
				return nil, errors.New("bencode: invalid list field")
			}
			if d.data[d.cursor] == 'e' {
				d.cursor += 1
				return list, nil
			}
			value, err := d.decode()
			if err != nil {
				return nil, err
			}
			list = append(list, value)
		}
	case 'd':
		d.cursor += 1
		dictionary := map[string]interface{}{}
		for {
			if d.cursor == d.length {
				return nil, errors.New("bencode: invalid dictionary field")
			}
			if d.data[d.cursor] == 'e' {
				d.cursor += 1
				return dictionary, nil
			}
			key, err := d.decodeBytes()
			if err != nil {
				return nil, errors.New("bencode: non-string dictionary key")
			}
			value, err := d.decode()
			if err != nil {
				return nil, err
			}
			dictionary[internal.B2S(key)] = value
		}
	default:
		return d.decodeBytes()
	}
}
