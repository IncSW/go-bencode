package decoder

import (
	"bytes"
	"errors"
	"strconv"
)

var (
	pow10i64 = [...]int64{
		1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
		1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	}
	pow10i64Len = len(pow10i64)
)

func (d *Decoder) parseInt(data []byte) (int64, error) {
	isNegative := false
	if data[0] == '-' {
		data = data[1:]
		isNegative = true
	}
	maxDigit := len(data)
	if maxDigit > pow10i64Len {
		return 0, errors.New("bencode: invalid length of number")
	}
	sum := int64(0)
	for i, b := range data {
		if b < '0' || b > '9' {
			return 0, errors.New("bencode: invalid integer byte: " + strconv.FormatUint(uint64(b), 10))
		}
		c := int64(b) - 48
		digitValue := pow10i64[maxDigit-i-1]
		sum += c * digitValue
	}
	if isNegative {
		return -1 * sum, nil
	}
	return sum, nil
}

func (d *Decoder) decodeInt() (interface{}, error) {
	d.cursor += 1
	index := bytes.IndexByte(d.data[d.cursor:], 'e')
	if index == -1 {
		return nil, errors.New("bencode: invalid integer field")
	}
	index += d.cursor
	integer, err := d.parseInt(d.data[d.cursor:index])
	if err != nil {
		return nil, err
	}
	d.cursor = index + 1
	return integer, nil
}
