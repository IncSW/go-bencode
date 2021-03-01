package bencode

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

// A Decoder reads and decodes Bencode values from an input stream.
type Decoder struct {
	r          io.Reader
	fromReader bool
	data       []byte
	length     int
	cursor     int
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r, fromReader: true}
}

// NewDecodeBytes returns a new decoder that decodes given bytes.
func NewDecodeBytes(src []byte) *Decoder {
	d := &Decoder{
		data:   src,
		length: len(src),
	}
	return d
}

// Decode writes the Bencode encoding of v to the stream.
func (d *Decoder) Decode(v interface{}) error {
	switch v.(type) {
	case nil:
		return errors.New("bencode: cannot marshal into nil")
	default:
		// we cannot catch other types, cause this might be a struct
		// so decoding below even if `v` isn't a supported type.
	}

	if d.fromReader {
		var err error
		d.data, err = ioutil.ReadAll(d.r)
		if err != nil {
			return fmt.Errorf("bencode: cannot read from reader: %w", err)
		}
		d.length = len(d.data)
	}
	if d.length == 0 {
		return errors.New("bencode: cannot decode empty input")
	}

	got, err := d.unmarshal()
	if err != nil {
		return fmt.Errorf("bencode: decode failed: %w", err)
	}
	return d.writeResult(v, got)
}

func (d *Decoder) writeResult(v, got interface{}) error {
	switch v := v.(type) {
	case *interface{}: // catch any type
		*v = got

	case *map[string]interface{}:
		val, ok := got.(map[string]interface{})
		if !ok {
			return fmt.Errorf("bencode: cannot decode %T into map", got)
		}
		*v = val

	case *[]interface{}:
		val, ok := got.([]interface{})
		if !ok {
			return fmt.Errorf("bencode: cannot decode %T into list", got)
		}
		*v = val

	default:
		// TODO: decode into struct if got map
	}
	return nil
}

func (d *Decoder) unmarshal() (interface{}, error) {
	switch d.data[d.cursor] {
	case 'i':
		return d.unmarshalInt()
	case 'd':
		return d.unmarshalMap()
	case 'l':
		return d.unmarshalList()
	default:
		return d.unmarshalString()
	}
}

func (d *Decoder) unmarshalInt() (int64, error) {
	d.cursor++
	index := bytes.IndexByte(d.data[d.cursor:], 'e')
	if index == -1 {
		return 0, errors.New("cannot process invalid integer")
	}

	index += d.cursor
	integer, err := strconv.ParseInt(b2s(d.data[d.cursor:index]), 10, 64)
	if err != nil {
		return 0, err
	}
	d.cursor = index + 1
	return integer, nil
}

func (d *Decoder) unmarshalMap() (interface{}, error) {
	dictionary := make(map[string]interface{})
	d.cursor++
	for {
		if d.cursor == d.length {
			return nil, errors.New("cannot process invalid dictionary")
		}
		if d.data[d.cursor] == 'e' {
			d.cursor++
			return dictionary, nil
		}

		// try to decode a dict key. it's a string by the Bencode specification
		key, err := d.unmarshalString()
		if err != nil {
			return nil, err
		}

		value, err := d.unmarshal()
		if err != nil {
			return nil, err
		}
		dictionary[b2s(key)] = value
	}
}

func (d *Decoder) unmarshalList() (interface{}, error) {
	list := make([]interface{}, 0)
	d.cursor++
	for {
		if d.cursor == d.length {
			return nil, errors.New("cannot process invalid list")
		}
		if d.data[d.cursor] == 'e' {
			d.cursor++
			return list, nil
		}
		value, err := d.unmarshal()
		if err != nil {
			return nil, err
		}
		list = append(list, value)
	}
}

func (d *Decoder) unmarshalString() ([]byte, error) {
	index := bytes.IndexByte(d.data[d.cursor:], ':')
	if index == -1 {
		return nil, errors.New("cannot process invalid string")
	}

	index += d.cursor
	strLen, err := strconv.ParseInt(b2s(d.data[d.cursor:index]), 10, 64)
	if err != nil {
		return nil, err
	}
	if strLen < 0 {
		return nil, errors.New("string length can not be a negative")
	}

	index++
	endIndex := index + int(strLen)
	if endIndex > d.length {
		return nil, errors.New("string length is not correct")
	}

	value := d.data[index:endIndex]
	d.cursor = endIndex
	return value, nil
}
