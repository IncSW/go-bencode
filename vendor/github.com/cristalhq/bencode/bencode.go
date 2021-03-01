package bencode

import (
	"bytes"
)

// Marshaler is the interface implemented by types that
// can marshal themselves into valid Bencode.
type Marshaler interface {
	MarshalBencode() ([]byte, error)
}

// Marshal returns bencode encoding of v.
func Marshal(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := NewEncoder(buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalTo returns bencode encoding of v written to dst.
func MarshalTo(dst []byte, v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(dst)
	enc := &Encoder{buf: buf}
	if err := enc.marshal(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unmarshaler is the interface implemented by types
// that can unmarshal a Bencode description of themselves.
type Unmarshaler interface {
	UnmarshalBencode([]byte) error
}

// Unmarshal parses the bencoded data and stores the result
// in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	d := NewDecodeBytes(data)
	if err := d.Decode(v); err != nil {
		return err
	}
	return nil
}
